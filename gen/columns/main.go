//go:build ignore

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

	outputPath, ok := os.LookupEnv("OUTPUT_PATH")
	if !ok {
		return errors.New("OUTPUT_PATH needed as environment variable")
	}

	credentialPath, ok := os.LookupEnv("CREDENTIAL_PATH")
	if !ok {
		return errors.New("CREDENTIAL_PATH needed as environment variable")
	}

	service, err := analyticsdata.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return err
	}

	response, err := service.Properties.GetMetadata("properties/0/metadata").Do()
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := genColumns(file, response); err != nil {
		return err
	}

	return nil
}

func genColumns(
	w io.Writer,
	response *analyticsdata.Metadata,
) error {
	if _, err := fmt.Fprintln(w, "package ga4data"); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "const ("); err != nil {
		return err
	}
	{
		for _, dimension := range response.Dimensions {
			title := strings.Title(dimension.ApiName)
			if strings.HasSuffix(title, "Id") {
				title = strings.TrimSuffix(title, "Id")
				title = title + "ID"
			}
			fmt.Fprintf(w, "// Dimension%s : %s\n", title, dimension.Description)
			fmt.Fprintf(w, "Dimension%s = %q\n", title, dimension.ApiName)
		}
		for _, metric := range response.Metrics {
			title := strings.Title(metric.ApiName)
			if strings.HasSuffix(title, "Id") {
				title = strings.TrimSuffix(title, "Id")
				title = title + "ID"
			}
			fmt.Fprintf(w, "// Metric%s : %s\n", title, metric.Description)
			fmt.Fprintf(w, "Metric%s = %q\n", title, metric.ApiName)
		}
	}
	if _, err := fmt.Fprintln(w, ")"); err != nil {
		return err
	}

	return nil
}
