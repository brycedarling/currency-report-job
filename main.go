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

	var output string
	flag.StringVar(&output, "output", "CSV", "CSV, HTML, XLS, or JSON")

	flag.Parse()

	if base == "" {
		base = "EUR"
	}

	if currencies == "" {
		log.Fatal(errors.New("missing currencies flag"))
	}

	output = strings.ToUpper(output)

	if !(output == "CSV" || output == "HTML" || output == "XLS" || output == "JSON") {
		log.Fatal(errors.New("output format must be CSV, HTML, XLS, or JSON"))
	}

	reports := fetchReports(currencies, base)

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
