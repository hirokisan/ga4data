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
dimensions := []string{"date"}
metrics := []string{"sessions"}
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
