package main

import (
	"context"
	"google.golang.org/api/sheets/v4"
	"log"
)

// Retrieve a token, saves the token, then returns the generated client.

func main() {

	srv, err := sheets.NewService(context.TODO())
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1HUIYUXbSX21ZgR93l0ryA-Q_PMoNGeIUlAIlHzT9dNc"
	readRange := "A:B"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	log.Println(resp.Values)

}

