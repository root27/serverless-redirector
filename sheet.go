package main

import (
	"context"
	"errors"
	"google.golang.org/api/sheets/v4"

	"log"
)

type sheet struct {
	Name string
	Id   string
}

func (s *sheet) Query() ([][]interface{}, error) {

	srv, err := sheets.NewService(context.TODO())

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	if s.Id == "" {

		log.Fatal("Google Sheets ID not specified")

		return nil, errors.New("Google Sheets ID not specified")

	}

	//NOTE: Set env variable
	spreadsheetId := s.Id
	readRange := "A:B"

	if s.Name != "" {

		readRange = s.Name + "!A:B"
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)

		return nil, err
	}

	log.Printf("Number of quering rows: %d\n", len(resp.Values))

	return resp.Values, nil
}
