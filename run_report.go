package ga4data

import (
	"context"

	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

const (
	// RunReportMaxRowsLimit : The API returns a maximum of 100,000 rows per request
	// ref: https://developers.google.com/analytics/devguides/reporting/data/v1/rest/v1beta/properties/runReport#body.request_body
	RunReportMaxRowsLimit = int64(100000)
)

// RunReportResponse :
type RunReportResponse struct {
	*analyticsdata.RunReportResponse

	RequestedCount int
}

// MergeResponse :
func (r *RunReportResponse) MergeResponse(resp *analyticsdata.RunReportResponse) {
	if r.RunReportResponse == nil {
		r.RunReportResponse = resp
		return
	}
	r.Rows = append(r.Rows, resp.Rows...)

	if quota := resp.PropertyQuota; quota != nil {
		r.PropertyQuota.ConcurrentRequests = quota.ConcurrentRequests
		r.PropertyQuota.PotentiallyThresholdedRequestsPerHour = quota.PotentiallyThresholdedRequestsPerHour
		r.PropertyQuota.ServerErrorsPerProjectPerHour = quota.ServerErrorsPerProjectPerHour

		if quota := quota.TokensPerDay; quota != nil {
			r.PropertyQuota.TokensPerDay.Consumed += quota.Consumed
			r.PropertyQuota.TokensPerDay.Remaining = quota.Remaining
		}
		if quota := quota.TokensPerHour; quota != nil {
			r.PropertyQuota.TokensPerHour.Consumed += quota.Consumed
			r.PropertyQuota.TokensPerHour.Remaining = quota.Remaining
		}
		if quota := quota.TokensPerProjectPerHour; quota != nil {
			r.PropertyQuota.TokensPerProjectPerHour.Consumed += quota.Consumed
			r.PropertyQuota.TokensPerProjectPerHour.Remaining = quota.Remaining
		}
	}
}

// RunReport :
func RunReport(
	ctx context.Context,
	service *analyticsdata.Service,
	propertyID string,
	request *analyticsdata.RunReportRequest,
) (*RunReportResponse, error) {
	var result RunReportResponse

	request.Limit = RunReportMaxRowsLimit
	for {
		request.Offset = RunReportMaxRowsLimit * int64(result.RequestedCount)

		resp, err := service.Properties.RunReport(propertyID, request).Do()
		if err != nil {
			return nil, err
		}

		result.MergeResponse(resp)
		result.RequestedCount++

		if resp.RowCount <= RunReportMaxRowsLimit*int64(result.RequestedCount) {
			break
		}
	}

	return &result, nil
}
