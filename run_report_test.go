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
	t.Run("quota consumed will be accumulated", func(t *testing.T) {
		wantRequestedCount := 2
		propertyQuota := analyticsdata.PropertyQuota{
			ConcurrentRequests: &analyticsdata.QuotaStatus{
				Consumed:  0,
				Remaining: 10,
			},
			PotentiallyThresholdedRequestsPerHour: &analyticsdata.QuotaStatus{
				Consumed:  0,
				Remaining: 120,
			},
			ServerErrorsPerProjectPerHour: &analyticsdata.QuotaStatus{
				Consumed:  0,
				Remaining: 10,
			},
			TokensPerDay: &analyticsdata.QuotaStatus{
				Consumed:  3,
				Remaining: 24997,
			},
			TokensPerHour: &analyticsdata.QuotaStatus{
				Consumed:  3,
				Remaining: 4997,
			},
			TokensPerProjectPerHour: &analyticsdata.QuotaStatus{
				Consumed:  3,
				Remaining: 1247,
			},
		}
		resBody := ga4datatest.RunReportResponseBody(
			ga4datatest.RunReportResponseBodyWithRowCount(int64(wantRequestedCount)*RunReportMaxRowsLimit),
			ga4datatest.RunReportResponseBodyWithQuota(propertyQuota),
		)
		testClient := ga4datatest.NewTestClient(resBody)

		service, err := analyticsdata.NewService(ctx, option.WithHTTPClient(testClient))
		require.NoError(t, err)

		request := CreateRunReportRequest(analyticsdata.DateRange{
			StartDate: "2022-10-01",
			EndDate:   "2022-10-01",
		},
			RunReportRequestWithPropertyQuota(),
		)

		resp, err := RunReport(ctx, service, "test", request)
		require.NoError(t, err)

		require.Equal(t, wantRequestedCount, resp.RequestedCount)

		assert.Equal(t, propertyQuota.TokensPerDay.Consumed*int64(wantRequestedCount), resp.PropertyQuota.TokensPerDay.Consumed)
		assert.Equal(t, propertyQuota.TokensPerHour.Consumed*int64(wantRequestedCount), resp.PropertyQuota.TokensPerHour.Consumed)
		assert.Equal(t, propertyQuota.TokensPerProjectPerHour.Consumed*int64(wantRequestedCount), resp.PropertyQuota.TokensPerProjectPerHour.Consumed)
	})
}
