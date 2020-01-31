package main

import (
	"html/template"
	"log"
	"os"
)

func writeHTMLReportFile(reports []RateReport) {
	file, err := os.Create("report.html")
	if err != nil {
		log.Fatal("cannot create file", err)
	}
	defer file.Close()

	t := template.Must(template.New("").Parse(`{{- range .}}<div>
	<h1>Currency: {{.Currency}}</h1>
	<ul>
	<li>Today Rate: {{.TodayRate}}</li>
	<li>One Day Ago Rate: {{.OneDayAgoRate}}</li>
	<li>OneDayAgoDelta: {{.OneDayAgoDelta}}</li>
	<li>SevenDaysAgoRate: {{.SevenDaysAgoRate}}</li>
	<li>SevenDaysAgoDelta: {{.SevenDaysAgoDelta}}</li>
	<li>OneMonthAgoRate: {{.OneMonthAgoRate}}</li>
	<li>OneMonthAgoDelta: {{.OneMonthAgoDelta}}</li>
	<li>OneYearAgoRate: {{.OneYearAgoRate}}</li>
	<li>OneYearAgoDelta: {{.OneYearAgoDelta}}</li>
	</ul>
</div>
{{end}}
`))

	err = t.Execute(file, reports)
	if err != nil {
		log.Fatal("cannot write to file", err)
	}
}
