package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Mahasiswa struct {
	Nama    string `json:"nama"`
	Umur    int    `json:"umur"`
	Jurusan string `json:"jurusan"`
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Hanya mendukung method POST!", http.StatusMethodNotAllowed)
		return
	}

	var mahasiswas []Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswas)
	if err != nil {
		// Log error untuk developer
		log.Println("Error decoding JSON:", err)
		// Kirim error ke user
		http.Error(w, "Gagal decode JSON. Pastikan format benar.", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"pesan":  "Berhasil membuat data!",
		"data":   mahasiswas,
		"status": http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set status code eksplisit
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", create)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
