package main

import (
	"log"
	"strings"
	"time"

	fixer "github.com/brycedarling/fixer-api-proxy/pkg"
)

type RateReport struct {
	Currency          string  `json:"currency"`
	TodayRate         float64 `json:"todayRate"`
	OneDayAgoRate     float64 `json:"oneDayAgoRate"`
	OneDayAgoDelta    float64 `json:"oneDayAgoDelta"`
	SevenDaysAgoRate  float64 `json:"sevenDaysAgoRate"`
	SevenDaysAgoDelta float64 `json:"sevenDaysAgoDelta"`
	OneMonthAgoRate   float64 `json:"oneMonthAgoRate"`
	OneMonthAgoDelta  float64 `json:"oneMonthAgoDelta"`
	OneYearAgoRate    float64 `json:"oneYearAgoRate"`
	OneYearAgoDelta   float64 `json:"oneYearAgoDelta"`
}

func fetchReports(currencies, base string) []RateReport {
	f, err := fixer.NewRateFetcher()
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	oneDayAgo := now.AddDate(0, 0, -1)
	sevenDaysAgo := now.AddDate(0, 0, -7)
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	var reports []RateReport

	layout := "2006-01-02"
	result := strings.Split(currencies, ",")
	for i := range result {
		currency := result[i]

		todayRate, err := f.FetchExchangeRate(now.Format(layout), base, currency)
		if err != nil {
			log.Fatal(err)
		}

		oneDayAgoRate, err := f.FetchExchangeRate(oneDayAgo.Format(layout), base, currency)
		if err != nil {
			log.Fatal(err)
		}

		sevenDaysAgoRate, err := f.FetchExchangeRate(sevenDaysAgo.Format(layout), base, currency)
		if err != nil {
			log.Fatal(err)
		}

		oneMonthAgoRate, err := f.FetchExchangeRate(oneMonthAgo.Format(layout), base, currency)
		if err != nil {
			log.Fatal(err)
		}

		oneYearAgoRate, err := f.FetchExchangeRate(oneYearAgo.Format(layout), base, currency)
		if err != nil {
			log.Fatal(err)
		}

		report := RateReport{
			Currency:          currency,
			TodayRate:         todayRate,
			OneDayAgoRate:     oneDayAgoRate,
			OneDayAgoDelta:    todayRate - oneDayAgoRate,
			SevenDaysAgoRate:  sevenDaysAgoRate,
			SevenDaysAgoDelta: todayRate - sevenDaysAgoRate,
			OneMonthAgoRate:   oneMonthAgoRate,
			OneMonthAgoDelta:  todayRate - oneMonthAgoRate,
			OneYearAgoRate:    oneYearAgoRate,
			OneYearAgoDelta:   todayRate - oneYearAgoRate,
		}

		reports = append(reports, report)
	}

	return reports
}
