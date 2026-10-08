# Combinations (Go TDD)

A small Go project that implements and tests a `combine` function. Given two
integers `n` and `k`, `combine` returns every possible combination of `k`
numbers chosen from the range `1..n`.

This project is primarily an exercise in **test-driven development**: the focus
is on exercising `combine` through table-driven unit tests, measuring coverage,
and benchmarking performance.

## The `combine` function

```go
func combine(n int, k int) ([][]int, error)
```

It generates all `k`-length combinations from `1..n` using a backtracking
helper (`backtrack`).

### Rules and return values

| Input                | Behavior                                      |
| -------------------- | --------------------------------------------- |
| Valid `n`, `k`       | Returns all combinations, `error` is `nil`    |
| `k == 0`             | Valid — returns an empty result, `error` is `nil` |
| `n <= 0`             | Invalid — returns a non-nil `error`           |
| `k < 0`              | Invalid — returns a non-nil `error`           |
| `k > n`              | Invalid — returns a non-nil `error`           |

Example: `combine(4, 2)` returns
`[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]`.

## Project layout

| File                   | Purpose                                      |
| ---------------------- | -------------------------------------------- |
| `combinations.go`      | Implementation of `combine` and `backtrack`  |
| `combinations_test.go` | Table-driven unit tests and a benchmark      |
| `go.mod`               | Go module definition (`combinations`)        |

## Requirements

- Go 1.27 or newer (see `go.mod`)

## Running the tests

Run the unit tests:

```sh
go test
```

Add `-v` for verbose, per-case output:

```sh
go test -v
```

### Coverage

Print a coverage summary in the terminal:

```sh
go test -cover
```

Generate a coverage profile and turn it into a browsable HTML report
(`coverage.html`):

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Open `coverage.html` in a browser to see which lines are covered.

### Benchmarks

Run the benchmarks:

```sh
go test -bench=.
```

Save the benchmark output to `benchmark.txt`:

```sh
go test -bench=. > benchmark.txt
```

> On a Unix-like shell you can use `go test -bench=. | tee benchmark.txt` instead.

## Generated reports

- `coverage.html` — HTML coverage report (produced by the coverage commands above)
- `benchmark.txt` — captured benchmark results (produced by the benchmark command above)
