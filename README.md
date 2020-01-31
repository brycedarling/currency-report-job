Set `FIXER_API_KEY` environment variable to your Fixer API key first before running:

```
export FIXER_API_KEY=changeme
```

Build program:

```
go build
```

Run program:

```
./currency-report-job -currencies USD,GBP
```

The `currencies` flag is a required comma separated list of currencies for the report job.

The `outputs` flag can be set to a comma separated list of `CSV`, `HTML`, or `JSON` - it defaults to all outputs `CSV,HTML,JSON`.

The `base` flag can be set for the base currency provided to the Fixer API, it defaults to `EUR`.

---

*TODO:* add XLS file output support

*TODO:* store reports in Google or Amazon storage on prod
