Set FIXER_API_KEY environment variable to your Fixer API key first before running:

```
export FIXER_API_KEY=changeme
```

Build program:

```
go build
```

Run program:

```
./currency-report-job -output HTML -currencies USD,GBP
```

The `output` flag can be set to `CSV, `HTML`, or `JSON`.

*TODO:* add XLS file output support

The `currencies` flag is a comma separated list of currencies for the report job.

The `base` flag can be set for the base currency provided to the Fixer API, it defaults to `EUR`.

---

*TODO:* I didn't get to adding XLS output support, or storing in Google or Amazon storage on prod.
