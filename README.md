# go-sqlite-benchmark

Benchmark of the SQLite3 database driver for Go.
Testcase is designed for my personal use cases.

# Insert 1000 rows into a table
```
BenchmarkEaton-16     1000 185098 ns/op
BenchmarkMattn_wal-16 1000 18053  ns/op
BenchmarkNcruses-16   1000 45907  ns/op
```
