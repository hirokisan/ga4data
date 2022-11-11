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
		if result.RequestedCount == 0 {
			result.RunReportResponse = resp
		} else {
			result.Rows = append(result.Rows, resp.Rows...)

			if request.ReturnPropertyQuota {
				result.PropertyQuota.ConcurrentRequests = resp.PropertyQuota.ConcurrentRequests
				result.PropertyQuota.PotentiallyThresholdedRequestsPerHour = resp.PropertyQuota.PotentiallyThresholdedRequestsPerHour
				result.PropertyQuota.ServerErrorsPerProjectPerHour = resp.PropertyQuota.ServerErrorsPerProjectPerHour
				result.PropertyQuota.TokensPerDay.Consumed += resp.PropertyQuota.TokensPerDay.Consumed
				result.PropertyQuota.TokensPerDay.Remaining = resp.PropertyQuota.TokensPerDay.Remaining
				result.PropertyQuota.TokensPerHour.Consumed += resp.PropertyQuota.TokensPerHour.Consumed
				result.PropertyQuota.TokensPerHour.Remaining = resp.PropertyQuota.TokensPerHour.Remaining
				result.PropertyQuota.TokensPerProjectPerHour.Consumed += resp.PropertyQuota.TokensPerProjectPerHour.Consumed
				result.PropertyQuota.TokensPerProjectPerHour.Remaining = resp.PropertyQuota.TokensPerProjectPerHour.Remaining
			}
		}

		result.RequestedCount++

		if resp.RowCount <= RunReportMaxRowsLimit*int64(result.RequestedCount) {
			break
		}
	}

	return &result, nil
}
