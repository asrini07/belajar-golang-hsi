package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	println("lib/pq berhasil diimport.")
	var _ *sql.DB 
}
