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

// RunReportRequestWithDimensionFilter :
func RunReportRequestWithDimensionFilter(filter *analyticsdata.FilterExpression) func(*analyticsdata.RunReportRequest) {
	return func(req *analyticsdata.RunReportRequest) {
		req.DimensionFilter = filter
	}
}

// RunReportRequestWithMetricFilter :
func RunReportRequestWithMetricFilter(filter *analyticsdata.FilterExpression) func(*analyticsdata.RunReportRequest) {
	return func(req *analyticsdata.RunReportRequest) {
		req.MetricFilter = filter
	}
}

// RunReportRequestWithPropertyQuota :
func RunReportRequestWithPropertyQuota() func(*analyticsdata.RunReportRequest) {
	return func(req *analyticsdata.RunReportRequest) {
		req.ReturnPropertyQuota = true
	}
}

// RunReportRequestWithDimensions :
func RunReportRequestWithDimensions(dimensions []string) func(*analyticsdata.RunReportRequest) {
	return func(req *analyticsdata.RunReportRequest) {
		req.Dimensions = Dimensions(dimensions)
	}
}

// RunReportRequestWithMetrics :
func RunReportRequestWithMetrics(metrics []string) func(*analyticsdata.RunReportRequest) {
	return func(req *analyticsdata.RunReportRequest) {
		req.Metrics = Metrics(metrics)
	}
}

// Dimensions :
func Dimensions(dimensions []string) []*analyticsdata.Dimension {
	dims := make([]*analyticsdata.Dimension, len(dimensions))
	for i := range dims {
		dims[i] = &analyticsdata.Dimension{
			Name: dimensions[i],
		}
	}
	return dims
}

// Metrics :
func Metrics(metrics []string) []*analyticsdata.Metric {
	mets := make([]*analyticsdata.Metric, len(metrics))
	for i := range mets {
		mets[i] = &analyticsdata.Metric{
			Name: metrics[i],
		}
	}
	return mets
}
