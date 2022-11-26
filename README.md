[![Go Report Card](https://goreportcard.com/badge/github.com/hirokisan/ga4data)](https://goreportcard.com/report/github.com/hirokisan/ga4data)
[![golangci-lint](https://github.com/hirokisan/bybit/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/hirokisan/bybit/actions/workflows/golangci-lint.yml)
[![test](https://github.com/hirokisan/bybit/actions/workflows/test.yml/badge.svg)](https://github.com/hirokisan/bybit/actions/workflows/test.yml)

# ga4data

ga4data is an client service wrapper of Google Analytics Data API (GA4) for the Go programming language.

## Usage

### RunReport

```
import (
	"github.com/hirokisan/ga4data"

	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

// prepare *analyticsdata.Service by yourself
// e.g. service, err := analyticsdata.NewService(ctx, option.WithHTTPClient(httpClient))

propertyID := "properties/xxx"
dimensions := []string{ga4data.DimensionDate}
metrics := []string{ga4data.MetricSessions}
response, err := ga4data.RunReport(ctx, service, propertyID, ga4data.CreateRunReportRequest(
	analyticsdata.DateRange{
		StartDate: "2022-10-01",
		EndDate: "2022-10-01",
	},
	ga4data.RunReportRequestWithDimensions(dimensions),
	ga4data.RunReportRequestWithMetrics(metrics),
))

// do as you want
```
