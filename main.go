package main

import (
	"context"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/url"
	"strings"
	"time"
)

type URLMap map[string]*url.URL

type cached struct {
	db        URLMap
	updatedAt time.Time
}

func (c *cached) refreshData() error {

}

func GetUrls(data [][]interface{}) URLMap {

	output := make(URLMap)

	for _, v := range data {

		key, ok := v[0].(string)

		if !ok {

			continue
		}

		key = strings.ToLower(key)

		value, ok := v[1].(string)

		if !ok {

			continue
		}

		url, err := url.Parse(value)

		if err != nil {

			continue
		}

		output[key] = url

	}

	log.Println(output)

	return output

}

func main() {

	srv, err := sheets.NewService(context.TODO())
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	//NOTE: Set env variable
	spreadsheetId := "1HUIYUXbSX21ZgR93l0ryA-Q_PMoNGeIUlAIlHzT9dNc"
	readRange := "A:B"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {

		log.Println("No data available")

		return
	}

}
