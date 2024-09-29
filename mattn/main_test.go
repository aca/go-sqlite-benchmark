package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
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

func BenchmarkMattn_wal(b *testing.B) {
	db, err := sql.Open("sqlite3", filepath.Join(dbdir, b.Name())+"?_journal_mode=WAL")
	if err != nil {
		b.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS journal (data TEXT);`)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := range b.N {
		_, err = db.Exec("INSERT INTO journal (data) VALUES (?)", ords[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

//	func BenchmarkZombizen(b *testing.B) {
//		conn, err := sqlite.OpenConn(filepath.Join(dbdir, b.Name()), sqlite.OpenReadWrite, sqlite.OpenCreate)
//		if err != nil {
//			b.Fatal(err)
//		}
//		err = sqlitex.ExecuteTransient(conn, schema, nil)
//		if err != nil {
//			b.Fatal(err)
//		}
//		defer conn.Close()
//
//		b.ResetTimer()
//		for i := range b.N {
//			err := sqlitex.Execute(conn, "INSERT INTO journal (data) VALUES (?)", &sqlitex.ExecOptions{
//				Args: []any{ords[i]},
//			})
//			if err != nil {
//				b.Fatal(err)
//			}
//		}
//	}
//
//	func BenchmarkZombizenWAL(b *testing.B) {
//		conn, err := sqlite.OpenConn(filepath.Join(dbdir, b.Name()), sqlite.OpenReadWrite, sqlite.OpenCreate, sqlite.OpenWAL)
//		if err != nil {
//			b.Fatal(err)
//		}
//		err = sqlitex.ExecuteTransient(conn, schema, nil)
//		if err != nil {
//			b.Fatal(err)
//		}
//		defer conn.Close()
//
//		b.ResetTimer()
//		for i := range b.N {
//			err := sqlitex.Execute(conn, "INSERT INTO journal (data) VALUES (?)", &sqlitex.ExecOptions{
//				Args: []any{ords[i]},
//			})
//			if err != nil {
//				b.Fatal(err)
//			}
//		}
//	}
//
//	func BenchmarkModerncWal(b *testing.B) {
//		db, err := sql.Open("sqlite", filepath.Join(dbdir, b.Name())+"?_journal=WAL&_timeout=5000")
//		if err != nil {
//			b.Fatal(err)
//		}
//		defer db.Close()
//
//		_, err = db.Exec("PRAGMA journal_mode=WAL")
//		if err != nil {
//			b.Fatal(err)
//		}
//
//		db.SetMaxOpenConns(1)
//		db.SetMaxIdleConns(1)
//		defer db.Close()
//
//		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS journal (data TEXT, processed INTEGER);`)
//		if err != nil {
//			panic(err)
//		}
//
//		b.ResetTimer()
//		for i := range b.N {
//			_, err = db.Exec("INSERT INTO journal (data) VALUES (?)", ords[i])
//			if err != nil {
//				panic(err)
//			}
//		}
//	}
//
//	func BenchmarkModerncNowal(b *testing.B) {
//		db, err := sql.Open("sqlite", filepath.Join(dbdir, b.Name())+"?_timeout=5000")
//		if err != nil {
//			panic(err)
//		}
//		defer db.Close()
//
//		// db.SetMaxOpenConns(1)
//		// db.SetMaxIdleConns(1)
//
//		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS journal (data TEXT, processed INTEGER);`)
//		if err != nil {
//			panic(err)
//		}
//
//		b.ResetTimer()
//		for i := range b.N {
//			_, err = db.Exec("INSERT INTO journal (data) VALUES (?)", ords[i])
//			if err != nil {
//				panic(err)
//			}
//		}
//	}

//
// func BenchmarkMattn_walwrc(b *testing.B) {
// 	db, err := sql.Open("sqlite3", filepath.Join(dbdir, b.Name())+"?mode=rwc&_journal=WAL&_timeout=5000&_fk=true")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
//
// 	_, err = db.Exec("PRAGMA journal_mode=WAL")
// 	if err != nil {
// 		b.Fatal(err)
// 	}
//
// 	db.SetMaxOpenConns(1)
// 	db.SetMaxIdleConns(1)
//
// 	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS journal (data TEXT, processed INTEGER);`)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	b.ResetTimer()
// 	for i := range b.N {
// 		_, err = db.Exec("INSERT INTO journal (data) VALUES (?)", ords[i])
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
//
// func BenchmarkMattn_nowal(b *testing.B) {
// 	db, err := sql.Open("sqlite3", filepath.Join(dbdir, b.Name()))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
//
// 	db.SetMaxOpenConns(1)
// 	db.SetMaxIdleConns(1)
//
// 	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS journal (data TEXT, processed INTEGER);`)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	b.ResetTimer()
// 	for i := range b.N {
// 		_, err = db.Exec("INSERT INTO journal (data) VALUES (?)", ords[i])
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
