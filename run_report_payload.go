package ga4data

import analyticsdata "google.golang.org/api/analyticsdata/v1beta"

// CreateRunReportRequest :
func CreateRunReportRequest(
	dr analyticsdata.DateRange,
	options ...func(req *analyticsdata.RunReportRequest),
) *analyticsdata.RunReportRequest {
	req := &analyticsdata.RunReportRequest{
		DateRanges: []*analyticsdata.DateRange{&dr},
	}
	for _, option := range options {
		option(req)
	}
	return req
}
