## Serverless URL Redirector


This is simple redirector service that can redirect a set of URL's in Google Sheets. You need to define [Google Sheets](https://docs.google.com/spreadsheets/d/14lESvPQuXoJfSHLk_gKs4FLYpW1vmWrtw8UbOw5CfmM/edit?usp=sharing) as shown below and deploy to
[Google Cloud](https://cloud.google.com/)

| Shortpath | Redirect |
| ---       | ---       |
| `gh`      | `https://github.com/root27`|
| `lin`     | `https://linkedin.com/in/ogzdo`|
| `ex`      | `https://docs.google.com/spreadsheets/d/14lESvPQuXoJfSHLk_gKs4FLYpW1vmWrtw8UbOw5CfmM/edit?usp=sharing`|

## How to Setup

1. Create new  [Google Sheet](https://sheets.new)

1. Set URL's of your desired shorpaths ([see example](https://docs.google.com/spreadsheets/d/14lESvPQuXoJfSHLk_gKs4FLYpW1vmWrtw8UbOw5CfmM/edit?usp=sharing))

1. Save ID of the your google sheet (example: 14lESvPQuXoJfSHLk_gKs4FLYpW1vmWrtw8UbOw5CfmM)

1. Click below button to deploy application to Cloud Run and provide sheet ID during deployment;

    [![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

1. Go to [Cloud Console](https://console.cloud.google.com/run) and click on `sheet-redirector` service. Copy the email address in `Service account`section.

1. Back to your Google Sheets and share it with this email address as "Viewer" access.

1. Enable the Google Sheets API in [here](https://console.developers.google.com/apis/api/sheets.googleapis.com/overview)

## Variable Config

You can configure several parameters in this service. The parameters are;

| Env. Variable | Description |
| ---           |   ---     |
| `SHEETNAME`(optional)   | If you want to manage multiple Google Sheets, you can provide Sheet name|
| `TTL` (optional)| You can set time how frequently the sheet must be queried (default: 5 seconds) |
`PORT` (optional) | Server port to listen (default: 8000)|



