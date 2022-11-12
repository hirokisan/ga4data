# columns

generate constant file for columns(dimensions,metrics)

## Preparation

see [document](https://developers.google.com/analytics/devguides/reporting/data/v1/quickstart-client-libraries) and do below.

- Create a project on GCP
- Download credential file of service account, and rename it as `credential.json`, place it in this directory
- Add service account email as GA4 Property user on GA4 dashboard

## Run

```console
$ make gen
```

