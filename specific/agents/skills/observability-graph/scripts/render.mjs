import { readFileSync, writeFileSync } from "node:fs";
import { createRequire } from "node:module";
import { dirname, join } from "node:path";
import { parseArgs } from "node:util";

import { Resvg } from "@resvg/resvg-js";
import * as vega from "vega";
import * as vl from "vega-lite";

const usage = `usage: node render.mjs (--data FILE | --spec FILE) [options]

  --data FILE      Prometheus/Loki matrix JSON, or an array of {time, series, value} rows
  --spec FILE      Vega-Lite spec to render with the house style instead of the default chart
  --out FILE       Output path ending in .png or .svg (default chart.png)
  --title TEXT     Chart title
  --subtitle TEXT  Subtitle, for the query, time range, and timezone
  --y-title TEXT   Y axis title, with the unit
  --y-format FMT   d3-format string for the Y axis, such as .0% or ~s
  --mark TYPE      line (default), area (stacked), or bar (stacked)
  --label NAME     Matrix label to name each series by
  --max-series N   Keep the N largest series (default 8); stacked charts sum the rest into "other"`;

const { values: opts } = parseArgs({
  options: {
    data: { type: "string" },
    spec: { type: "string" },
    out: { type: "string", default: "chart.png" },
    title: { type: "string" },
    subtitle: { type: "string" },
    "y-title": { type: "string" },
    "y-format": { type: "string" },
    mark: { type: "string", default: "line" },
    label: { type: "string" },
    "max-series": { type: "string", default: "8" },
    width: { type: "string", default: "1000" },
    height: { type: "string", default: "450" },
    help: { type: "boolean", short: "h" },
  },
});

if (opts.help) {
  console.log(usage);
  process.exit(0);
}
if (!opts.data && !opts.spec) {
  console.error(usage);
  process.exit(1);
}

const font = "Inter";
const require = createRequire(import.meta.url);
const fontDir = dirname(require.resolve("@expo-google-fonts/inter/package.json"));
const fontFiles = ["400Regular/Inter_400Regular.ttf", "600SemiBold/Inter_600SemiBold.ttf"].map((f) => join(fontDir, f));

const config = {
  background: "#ffffff",
  padding: 24,
  font,
  view: { stroke: null },
  title: { fontSize: 22, fontWeight: 600, anchor: "start", subtitleFontSize: 14, subtitleColor: "#555", subtitlePadding: 6, offset: 16 },
  axis: { labelFontSize: 13, titleFontSize: 14, titleFontWeight: 600, labelColor: "#333", titleColor: "#333", gridColor: "#e6e6e6", domainColor: "#999", tickColor: "#999" },
  axisX: { grid: false, labelAngle: 0, labelFlush: true, labelOverlap: true },
  axisY: { domain: false, ticks: false, tickCount: 6, labelPadding: 8 },
  legend: { orient: "bottom", direction: "horizontal", labelFontSize: 13, symbolStrokeWidth: 3, symbolSize: 200, columns: 4, labelLimit: 320 },
  line: { strokeWidth: 2.5, interpolate: "monotone" },
  area: { opacity: 0.85, interpolate: "monotone" },
  range: { category: { scheme: "tableau10" } },
};

function seriesName(metric) {
  if (opts.label && metric[opts.label] !== undefined) return String(metric[opts.label]);
  const labels = Object.entries(metric).filter(([k]) => k !== "__name__");
  if (labels.length === 0) return metric.__name__ ?? "value";
  return labels.map(([k, v]) => (labels.length === 1 ? v : `${k}=${v}`)).join(", ");
}

function toMillis(t) {
  const ms = typeof t === "number" ? (t < 1e11 ? t * 1000 : t) : /^\d+(\.\d+)?$/.test(t) ? toMillis(Number(t)) : Date.parse(t);
  if (!Number.isFinite(ms)) throw new Error(`invalid time: ${JSON.stringify(t)}`);
  return ms;
}

function toRows(input) {
  if (Array.isArray(input)) {
    return input.map((r) => ({ time: toMillis(r.time), series: String(r.series ?? "value"), value: toValue(r.value) }));
  }
  const matrix = Array.isArray(input.data) ? input.data : input.data?.result;
  if (!Array.isArray(matrix)) throw new Error("--data must be an array of {time, series, value} rows or a Prometheus/Loki matrix ({data: {result: [...]}})");
  return matrix.flatMap((s) => {
    const series = seriesName(s.metric ?? {});
    const points = s.values ?? (s.value ? [s.value] : []);
    return points.map(([t, v]) => ({ time: toMillis(Number(t)), series, value: toValue(v) }));
  });
}

function toValue(v) {
  const n = v === null || v === undefined || v === "" ? NaN : Number(v);
  return Number.isFinite(n) ? n : null;
}

function limitSeries(rows, stacked) {
  const max = Number(opts["max-series"]);
  const totals = new Map();
  for (const r of rows) totals.set(r.series, (totals.get(r.series) ?? 0) + Math.abs(r.value ?? 0));
  if (totals.size <= max) return rows;
  const keepCount = stacked ? max - 1 : max;
  const keep = new Set([...totals].sort((a, b) => b[1] - a[1]).slice(0, keepCount).map(([s]) => s));
  const dropped = totals.size - keep.size;
  if (!stacked) {
    console.error(`render: kept the ${keep.size} largest of ${totals.size} series; dropped ${dropped}`);
    return rows.filter((r) => keep.has(r.series));
  }
  const other = new Map();
  for (const r of rows) {
    if (keep.has(r.series) || r.value === null) continue;
    other.set(r.time, (other.get(r.time) ?? 0) + r.value);
  }
  console.error(`render: kept the ${keep.size} largest of ${totals.size} series; summed ${dropped} into "other"`);
  return [...rows.filter((r) => keep.has(r.series)), ...[...other].map(([time, value]) => ({ time, series: "other", value }))];
}

function timeFormat(rows) {
  const times = rows.map((r) => new Date(r.time).getTime());
  const spanHours = (Math.max(...times) - Math.min(...times)) / 3.6e6;
  if (spanHours <= 36) return "%H:%M";
  if (spanHours <= 24 * 7) return "%a %d %H:%M";
  return "%b %d";
}

function defaultSpec(input) {
  if (!["line", "area", "bar"].includes(opts.mark)) throw new Error(`--mark must be line, area, or bar, got ${opts.mark}`);
  const stacked = opts.mark === "area" || opts.mark === "bar";
  const rows = limitSeries(input, stacked);
  if (!rows.some((r) => r.value !== null)) throw new Error("no data points to plot");
  const counts = new Map();
  for (const r of rows) if (r.value !== null) counts.set(r.series, (counts.get(r.series) ?? 0) + 1);
  const sparse = Math.min(...counts.values()) <= 60;
  const multi = counts.size > 1;
  const peaks = new Map();
  for (const r of rows) if (r.value !== null) peaks.set(r.series, Math.max(peaks.get(r.series) ?? -Infinity, r.value));
  const order = [...peaks].sort((a, b) => (a[0] === "other") - (b[0] === "other") || b[1] - a[1]).map(([s]) => s);
  return {
    width: Number(opts.width),
    height: Number(opts.height),
    data: { values: rows },
    mark: { type: opts.mark, ...(opts.mark === "line" && sparse ? { point: { size: 36 } } : {}) },
    encoding: {
      x: { field: "time", type: "temporal", title: null, axis: { format: timeFormat(rows) } },
      y: {
        field: "value",
        type: "quantitative",
        title: opts["y-title"] ?? null,
        stack: stacked ? "zero" : null,
        scale: { zero: true, nice: true },
        ...(opts["y-format"] ? { axis: { format: opts["y-format"] } } : {}),
      },
      ...(multi ? { color: { field: "series", type: "nominal", title: null, sort: order } } : {}),
    },
  };
}

const spec = opts.spec ? JSON.parse(readFileSync(opts.spec, "utf8")) : defaultSpec(toRows(JSON.parse(readFileSync(opts.data, "utf8"))));
if (opts.title) spec.title = { text: opts.title, ...(opts.subtitle ? { subtitle: opts.subtitle } : {}) };
spec.config = { ...config, ...(spec.config ?? {}) };

const view = new vega.View(vega.parse(vl.compile(spec).spec), { renderer: "none" });
const svg = await view.toSVG();
view.finalize();

if (opts.out.endsWith(".svg")) {
  writeFileSync(opts.out, svg);
} else {
  const png = new Resvg(svg, {
    fitTo: { mode: "zoom", value: 2 },
    font: { fontFiles, loadSystemFonts: false, defaultFontFamily: font },
  }).render().asPng();
  writeFileSync(opts.out, png);
}
console.log(opts.out);
