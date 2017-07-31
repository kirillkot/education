package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func needPin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("pin") != "9999" {
			http.Error(w, "wrong pin", http.StatusForbidden)
			return
		}

		h(w, r)
	}
}

func main() {
	http.HandleFunc("/", needPin(handler))
	http.ListenAndServe(":1234", nil)
}
