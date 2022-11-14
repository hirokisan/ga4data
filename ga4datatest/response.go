package ga4datatest

import (
	"encoding/json"

	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

// RunReportResponseBody :
func RunReportResponseBody(options ...func(*analyticsdata.RunReportResponse)) []byte {
	respBody := &analyticsdata.RunReportResponse{
		DimensionHeaders: []*analyticsdata.DimensionHeader{},
		MetricHeaders:    []*analyticsdata.MetricHeader{},
		Rows:             []*analyticsdata.Row{},
		RowCount:         int64(0),
		PropertyQuota:    &analyticsdata.PropertyQuota{},
	}
	for _, option := range options {
		option(respBody)
	}
	bytesBody, err := json.Marshal(respBody)
	if err != nil {
		panic(err)
	}
	return bytesBody
}

// RunReportResponseBodyWithRows :
func RunReportResponseBodyWithRows(
	dimensionHeaders []string,
	metricHeaders []string,
	dimensionValues [][]string,
	metricValues [][]string,
) func(*analyticsdata.RunReportResponse) {
	dimHeaders := make([]*analyticsdata.DimensionHeader, len(dimensionHeaders))
	for i := range dimHeaders {
		dimHeaders[i] = &analyticsdata.DimensionHeader{
			Name: dimensionHeaders[i],
		}
	}
	metHeaders := make([]*analyticsdata.MetricHeader, len(metricHeaders))
	for i := range metHeaders {
		metHeaders[i] = &analyticsdata.MetricHeader{
			Name: metricHeaders[i],
		}
	}
	if len(dimensionValues) != len(metricValues) {
		panic("dimensionValues and metricValues should be same length")
	}
	rows := make([]*analyticsdata.Row, len(dimensionValues))
	for i := range rows {
		dimValues := make([]*analyticsdata.DimensionValue, len(dimensionValues[i]))
		for j := range dimValues {
			dimValues[j] = &analyticsdata.DimensionValue{
				Value: dimensionValues[i][j],
			}
		}
		metValues := make([]*analyticsdata.MetricValue, len(metricValues[i]))
		for j := range metValues {
			metValues[j] = &analyticsdata.MetricValue{
				Value: metricValues[i][j],
			}
		}
		rows[i] = &analyticsdata.Row{
			DimensionValues: dimValues,
			MetricValues:    metValues,
		}
	}
	return func(res *analyticsdata.RunReportResponse) {
		res.DimensionHeaders = dimHeaders
		res.MetricHeaders = metHeaders
		res.Rows = rows
		res.RowCount = int64(len(rows))
	}
}

// RunReportResponseBodyWithRowCount :
func RunReportResponseBodyWithRowCount(count int64) func(*analyticsdata.RunReportResponse) {
	return func(res *analyticsdata.RunReportResponse) {
		res.RowCount = count
	}
}
