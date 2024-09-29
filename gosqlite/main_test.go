package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/eatonphil/gosqlite"
)

var ords []string
var dbdir = "../"

func TestMain(m *testing.M) {
	var ord = strings.Repeat("0", 374-8)
	for i := 0; i < 10000; i++ {
		ords = append(ords, ord+fmt.Sprintf("%8d", i))
	}

	os.Exit(m.Run())
}

func BenchmarkEaton(b *testing.B) {
	conn, err := gosqlite.Open(filepath.Join(dbdir, b.Name()))
	if err != nil {
		b.Fatal(err)
	}

	err = conn.Exec(`PRAGMA journal_mode=WAL`)
	if err != nil {
		b.Fatal(err)
	}

	err = conn.Exec(`CREATE TABLE IF NOT EXISTS journal (data TEXT);`)
	if err != nil {
		b.Fatal(err)
	}

	stmt, err := conn.Prepare(`INSERT INTO journal (data) VALUES (?)`)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := range b.N {
		err = stmt.Exec(ords[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}
