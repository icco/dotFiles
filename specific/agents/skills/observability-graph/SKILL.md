---
name: observability-graph
description: Use when the user wants a graph, chart, or plot of metrics, logs, or traces from Grafana, Prometheus, Loki, Datadog, or another observability provider, especially to share it, upload it to photos.natwelch.com, or show it inline in Markdown.
---

# Observability graph

Fetch the series to a file, render a PNG with the bundled renderer, look at it, upload it to photos.natwelch.com, and reply with the inline image and its link.

## Setup

`<skill>` is this skill's base directory. Check the renderer:

```sh
node <skill>/scripts/render.mjs --help
```

If that fails with `ERR_MODULE_NOT_FOUND`, install the pinned dependencies next to the real file (symlinked skills resolve to their source checkout):

```sh
npm ci --prefix "$(dirname "$(realpath <skill>/scripts/render.mjs)")"
```

- opencode-server Docker image: already installed.
- macOS: `brew install node`, then the `npm ci` above once. The font ships with the dependencies, so no system fonts or browser are needed.

## 1. Fetch the data to a file

Write results straight to a file in a `mktemp -d` directory. Do not pull long series through an MCP tool into the conversation.

Grafana (`GRAFANA_URL`, `GRAFANA_SERVICE_ACCOUNT_TOKEN`) proxies every datasource. List UIDs with `curl -sS -H "Authorization: Bearer $GRAFANA_SERVICE_ACCOUNT_TOKEN" "${GRAFANA_URL%/}/api/datasources" | jq -r '.[] | "\(.uid) \(.type) \(.name)"'`.

```sh
dir=$(mktemp -d)
end=$(date +%s); start=$((end - 6*3600)); step=$(( (end - start) / 300 ))
curl -sS --fail-with-body -G \
  -H "Authorization: Bearer $GRAFANA_SERVICE_ACCOUNT_TOKEN" \
  "${GRAFANA_URL%/}/api/datasources/proxy/uid/prometheus/api/v1/query_range" \
  --data-urlencode 'query=sum by (instance) (rate(node_cpu_seconds_total{mode!="idle"}[5m]))' \
  --data-urlencode "start=$start" --data-urlencode "end=$end" --data-urlencode "step=$step" \
  -o "$dir/data.json"
jq -c '{status, series: (.data.result | length), labels: [.data.result[].metric]}' "$dir/data.json"
```

| Source | Endpoint | Pass to `--data` |
|---|---|---|
| Prometheus via Grafana | `/api/datasources/proxy/uid/<uid>/api/v1/query_range` | as is |
| Loki via Grafana | `/api/datasources/proxy/uid/<uid>/loki/api/v1/query_range` with a metric query, such as `sum by (service_name) (count_over_time({service_name=~".+"}[5m]))` | as is |
| Prometheus direct | `<prom>/api/v1/query_range` | as is |
| Datadog, CloudWatch, Tempo, CSV, anything else | the provider's API or CLI | `jq` to `[{"time": <epoch s, epoch ms, or ISO>, "series": "name", "value": 1.2}]` |

Aim for 200 to 500 points per series (`step = range / 300`). If only an MCP tool can reach the source, reduce the result to compact rows inside `execute` and save that with the write tool.

## 2. Shape it for reading

- Eight series at most. In a range query, `topk` picks per step and returns more than k series. Rank with an instant query over the whole range, such as `topk(5, sum by (x) (increase(m[6h])))`, then graph those with a `=~` regex. The renderer otherwise keeps the 8 largest and says so on stderr: line charts drop the rest, and stacked charts sum them into "other". Say what was dropped in the subtitle.
- Convert to human units in the query (`/ 2^30` for GiB) or with `--y-format` (`.0%` for ratios, `~s` for SI), and name the unit in `--y-title`.
- `line` for rates and gauges, `area` (stacked) for parts of a whole, `bar` (stacked) for counts per bucket.
- The title states what the reader learns ("Memory used per host"), not the metric name. The subtitle holds the source or query, the range, and the timezone.
- Times render in the process timezone. Prefix `TZ=America/New_York` (or the user's zone) when the user is not on UTC, and name the zone in the subtitle.
- Null, `NaN`, and `Inf` samples become gaps, not zeros. Series with few points get point markers so isolated samples stay visible.

## 3. Render and look

```sh
node <skill>/scripts/render.mjs --data "$dir/data.json" --out "$dir/chart.png" \
  --title "CPU usage by host" --subtitle "non-idle cores, rate 5m, last 6h, UTC" --y-title "busy cores"
```

See `--help` for `--mark`, `--label`, `--y-format`, `--max-series`, and size. For anything else, write a Vega-Lite spec and pass `--spec spec.json`; the house style still applies.

Read the PNG before uploading. Re-render if the legend is truncated, one series flattens the rest (split the chart), or a series is missing.

## 4. Upload

The upload is public. Host, service, and job names are fine. Ask before uploading a chart whose labels or text contain tokens, credentials, email addresses, customer names, or other personal data.

```sh
curl -sS --fail-with-body -F "photo=@$dir/chart.png" https://photos.natwelch.com/api/upload | jq -r '.files[0].url'
```

The response is `{"success": true, "files": [{"path": "...", "url": "https://images.natwelch.com/photos/<year>/<id>.png"}]}`. An empty `files` list means the extension was rejected.

## 5. Reply

Show the image inline, then the bare link:

```markdown
![CPU usage by host, last 6h](https://images.natwelch.com/photos/2026/0S19ZB8XYFMH7.png)

https://images.natwelch.com/photos/2026/0S19ZB8XYFMH7.png
```

When the data came from Grafana, add an Explore link:

```sh
printf '%s/explore?schemaVersion=1&panes=%s\n' "${GRAFANA_URL%/}" "$(jq -rn --arg q "$query" --arg ds prometheus \
  '{a: {datasource: $ds, queries: [{refId: "A", expr: $q, datasource: {uid: $ds}}], range: {from: "now-6h", to: "now"}}} | tojson | @uri')"
```

Delete the `mktemp` directory unless the user wants the files.
