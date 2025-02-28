package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		fmt.Fprintln(w, id)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		id := r.FormValue("id")
		nama := r.FormValue("nama")
		umur := r.FormValue("umur")

		fmt.Fprintln(w, id)
		fmt.Fprintln(w, nama)
		fmt.Fprintln(w, umur)
	}
}

func responseJson(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"pesan":  "berhasil menambah data !",
		"status": http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/search", search)
	mux.HandleFunc("/", create)
	mux.HandleFunc("/response-json", responseJson)

	http.ListenAndServe(":8080", mux)
}
