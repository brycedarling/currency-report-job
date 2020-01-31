package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeCSVReportFile(reports []RateReport) {
	file, err := os.Create("report.csv")
	if err != nil {
		log.Fatal("cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Currency", "TodayRate", "OneDayAgoRate", "OneDayAgoDelta", "SevenDaysAgoRate", "SevenDaysAgoDelta", "OneMonthAgoRate", "OneMonthAgoDelta", "OneYearAgoRate", "OneYearAgoDelta"})

	for _, report := range reports {
		err := writer.Write([]string{
			report.Currency,
			fmt.Sprintf("%f", report.TodayRate),
			fmt.Sprintf("%f", report.OneDayAgoRate),
			fmt.Sprintf("%f", report.OneDayAgoDelta),
			fmt.Sprintf("%f", report.SevenDaysAgoRate),
			fmt.Sprintf("%f", report.SevenDaysAgoDelta),
			fmt.Sprintf("%f", report.OneMonthAgoRate),
			fmt.Sprintf("%f", report.OneMonthAgoDelta),
			fmt.Sprintf("%f", report.OneYearAgoRate),
			fmt.Sprintf("%f", report.OneYearAgoDelta),
		})
		if err != nil {
			log.Fatal("cannot write to file", err)
		}
	}
}
