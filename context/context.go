package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	err error
	db  *sql.DB
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func needPin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pin string
		if err := db.QueryRow("SELECT pin FROM pins").Scan(&pin); err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		if r.FormValue("pin") != pin {
			http.Error(w, "wrong pin", http.StatusForbidden)
			return
		}

		h(w, r)
	}
}

func main() {
	if db, err = sql.Open("sqlite3", "./pins.db"); err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", needPin(handler))
	http.ListenAndServe(":1234", nil)
}
