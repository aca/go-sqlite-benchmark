# go-sqlite-benchmark

Benchmark of the SQLite3 database driver for Go.
Testcase is designed for my personal use cases.

### Insert 1000 rows into a table

```
goos: linux
goarch: amd64
pkg: drivers
cpu: AMD Ryzen 7 8845HS w/ Radeon 780M Graphics

journal_mode=WAL

BenchmarkEaton-16     1000 403731 ns/op
BenchmarkMattn_wal-16 1000 19675  ns/op
BenchmarkNcruses-16   1000 53337  ns/op
```
