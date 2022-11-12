//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hirokisan/ga4data"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	propertyID, ok := os.LookupEnv("PROPERTY_ID")
	if !ok {
		return errors.New("PROPERTY_ID needed as environment variable")
	}
	credentialPath, ok := os.LookupEnv("CREDENTIAL_PATH")
	if !ok {
		return errors.New("CREDENTIAL_PATH needed as environment variable")
	}

	service, err := analyticsdata.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return err
	}

	dimensions := []string{"date"}
	metrics := []string{"sessions"}
	response, err := ga4data.RunReport(ctx, service, propertyID, ga4data.CreateRunReportRequest(
		analyticsdata.DateRange{
			StartDate: "2022-10-01",
			EndDate:   "2022-10-01",
		},
		ga4data.RunReportRequestWithDimensions(dimensions),
		ga4data.RunReportRequestWithMetrics(metrics),
		ga4data.RunReportRequestWithPropertyQuota(),
	))
	if err != nil {
		return err
	}

	printResponse(response)

	return nil
}

func printResponse(response *ga4data.RunReportResponse) {
	fmt.Println("==== dimensionHeaders ====")
	for _, header := range response.DimensionHeaders {
		fmt.Println(header.Name)
	}

	fmt.Println("==== metricHeaders ====")
	for _, header := range response.MetricHeaders {
		fmt.Println(header.Name)
	}

	fmt.Println("==== propertyQuota ====")
	fmt.Printf("concurrentRequests: consumed %d, remaining %d\n",
		response.PropertyQuota.ConcurrentRequests.Consumed,
		response.PropertyQuota.ConcurrentRequests.Remaining,
	)
	fmt.Printf("potentiallyThresholdedRequestsPerHour: consumed %d, remaining %d\n",
		response.PropertyQuota.PotentiallyThresholdedRequestsPerHour.Consumed,
		response.PropertyQuota.PotentiallyThresholdedRequestsPerHour.Remaining,
	)
	fmt.Printf("serverErrorsPerProjectPerHour: consumed %d, remaining %d\n",
		response.PropertyQuota.ServerErrorsPerProjectPerHour.Consumed,
		response.PropertyQuota.ServerErrorsPerProjectPerHour.Remaining,
	)
	fmt.Printf("tokensPerDay: consumed %d, remaining %d\n",
		response.PropertyQuota.TokensPerDay.Consumed,
		response.PropertyQuota.TokensPerDay.Remaining,
	)
	fmt.Printf("tokensPerHour: consumed %d, remaining %d\n",
		response.PropertyQuota.TokensPerHour.Consumed,
		response.PropertyQuota.TokensPerHour.Remaining,
	)
	fmt.Printf("tokensPerProjectPerHour: consumed %d, remaining %d\n",
		response.PropertyQuota.TokensPerProjectPerHour.Consumed,
		response.PropertyQuota.TokensPerProjectPerHour.Remaining,
	)

	fmt.Println("==== rowCount ====")
	fmt.Println(response.RowCount)

	fmt.Println("==== rows ====")
	for _, row := range response.Rows {
		var dimensions []string
		for _, dim := range row.DimensionValues {
			dimensions = append(dimensions, dim.Value)
		}
		var metrics []string
		for _, met := range row.MetricValues {
			metrics = append(metrics, met.Value)
		}
		fmt.Printf("dimensions: %s, metrics: %s \n", dimensions, metrics)
	}
}
