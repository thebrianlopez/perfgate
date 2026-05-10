# perfgate

Statistical performance gating for shell commands. Benchmark any command, save a baseline, and enforce regression thresholds in CI.

## Install

```sh
go install github.com/thebrianlopez/perfgate/cmd/perfgate@latest
```

## Usage

### 1. Capture a baseline

```sh
perfgate run --cmd "my-tool --flag" --runs 20 --warmup 3 --save baseline.json
```

Output:

```
Results (20 runs):
  Mean:   0.1234s
  Median: 0.1201s
  P95:    0.1489s
  StdDev: 0.0091s
  Min:    0.1187s
  Max:    0.1521s
```

### 2. Compare two baselines

```sh
perfgate compare --before baseline.json --after current.json
```

Output:

```
Before (20 runs):
  Mean:   0.1234s
  ...

After (20 runs):
  Mean:   0.1301s
  ...

Change: +5.43%
```

### 3. Run a gate (pass/fail for CI)

```sh
perfgate gate \
  --cmd "my-tool --flag" \
  --baseline baseline.json \
  --max-regression 5.0 \
  --runs 20
```

Exits 0 on pass, 1 on regression. Use `--format json` for machine-readable output.

## Flags

### `run`

| Flag | Default | Description |
|------|---------|-------------|
| `--cmd` | required | Shell command to benchmark |
| `--runs` | 10 | Number of timed iterations |
| `--warmup` | 2 | Warmup iterations (discarded) |
| `--save` | | Save results to JSON file |
| `--format` | text | Output format: `text` or `json` |

### `compare`

| Flag | Default | Description |
|------|---------|-------------|
| `--before` | required | Baseline results file |
| `--after` | required | Current results file |
| `--format` | text | Output format: `text` or `json` |

### `gate`

| Flag | Default | Description |
|------|---------|-------------|
| `--cmd` | required | Shell command to benchmark |
| `--baseline` | required | Baseline results file |
| `--runs` | 10 | Number of timed iterations |
| `--warmup` | 2 | Warmup iterations (discarded) |
| `--max-regression` | 5.0 | Max allowed regression % before FAIL |
| `--min-improvement` | 0 | Min required improvement % (optional) |
| `--format` | text | Output format: `text` or `json` |

## How it works

`perfgate run` executes your command N times (plus warmup), measures wall-clock duration for each run, and computes mean, median, P95, stddev, min, and max. Results are saved as JSON.

`perfgate gate` runs the same benchmark against a saved baseline and computes percent change in mean duration. If the change exceeds `--max-regression`, the process exits 1 — failing the CI step.

## Stats computed

- **Mean** — average wall-clock time across all runs
- **Median** — middle value (robust to outliers)
- **P95** — 95th percentile (catches tail latency)
- **StdDev** — spread of measurements
- **Min / Max** — range

## License

MIT
