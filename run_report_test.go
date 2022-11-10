package ga4data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

func TestRunReport(t *testing.T) {
	ctx := context.Background()

	service, err := analyticsdata.NewService(ctx)
	require.NoError(t, err)

	request := CreateRunReportRequest(analyticsdata.DateRange{
		StartDate: "2022-10-01",
		EndDate:   "2022-10-01",
	})

	_, err = RunReport(ctx, service, "test", request)
	require.NoError(t, err)
}
