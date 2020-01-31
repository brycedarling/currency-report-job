package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	apiKey := os.Getenv("FIXER_API_KEY")
	if apiKey == "" {
		log.Fatal(errors.New("missing FIXER_API_KEY env var"))
	}

	var base string
	flag.StringVar(&base, "base", "", "base currency for conversion")

	var currencies string
	flag.StringVar(&currencies, "currencies", "", "currencies for report job")

	var outputs string
	flag.StringVar(&outputs, "outputs", "CSV,HTML,JSON", "CSV, HTML, or JSON")

	flag.Parse()

	if base == "" {
		base = "EUR"
	}

	if currencies == "" {
		log.Fatal(errors.New("missing currencies flag"))
	}

	reports := fetchReports(currencies, base)

	result := strings.Split(strings.ToUpper(outputs), ",")
	for i := range result {
		output := strings.TrimSpace(result[i])

		if !(output == "CSV" || output == "HTML" || output == "JSON") {
			log.Fatal(errors.New("output format must be CSV, HTML, or JSON"))
		}

		switch output {
		case "CSV":
			writeCSVReportFile(reports)
		case "HTML":
			writeHTMLReportFile(reports)
		case "JSON":
			writeJSONReportFile(reports)
		case "XLS":
			log.Fatal("xls is unsupported")
		}
	}
}
