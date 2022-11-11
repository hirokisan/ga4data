package ga4data

import (
	"context"
	"testing"

	"github.com/hirokisan/ga4data/ga4datatest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

func TestRunReport(t *testing.T) {
	ctx := context.Background()

	t.Run("no paging", func(t *testing.T) {
		wantRequestedCount := 1
		resBody := ga4datatest.RunReportResponseBody(
			ga4datatest.RunReportResponseBodyWithRowCount(int64(wantRequestedCount) * RunReportMaxRowsLimit),
		)
		testClient := ga4datatest.NewTestClient(resBody)

		service, err := analyticsdata.NewService(ctx, option.WithHTTPClient(testClient))
		require.NoError(t, err)

		request := CreateRunReportRequest(analyticsdata.DateRange{
			StartDate: "2022-10-01",
			EndDate:   "2022-10-01",
		})

		resp, err := RunReport(ctx, service, "test", request)
		require.NoError(t, err)

		assert.Equal(t, wantRequestedCount, resp.RequestedCount)
	})
	t.Run("paging all", func(t *testing.T) {
		wantRequestedCount := 2
		resBody := ga4datatest.RunReportResponseBody(
			ga4datatest.RunReportResponseBodyWithRowCount(int64(wantRequestedCount) * RunReportMaxRowsLimit),
		)
		testClient := ga4datatest.NewTestClient(resBody)

		service, err := analyticsdata.NewService(ctx, option.WithHTTPClient(testClient))
		require.NoError(t, err)

		request := CreateRunReportRequest(analyticsdata.DateRange{
			StartDate: "2022-10-01",
			EndDate:   "2022-10-01",
		})

		resp, err := RunReport(ctx, service, "test", request)
		require.NoError(t, err)

		assert.Equal(t, wantRequestedCount, resp.RequestedCount)
	})
}
