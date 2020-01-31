package main

import (
	"encoding/json"
	"log"
	"os"
)

func writeJSONReportFile(reports []RateReport) {
	file, err := os.Create("report.json")
	if err != nil {
		log.Fatal("cannot create file", err)
	}
	defer file.Close()

	jsonBytes, err := json.Marshal(reports)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		log.Fatal("cannot write to file", err)
	}
}
