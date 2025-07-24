package main

import (
	"fmt"
	"net/http"
	 "strconv"
	"encoding/json"
	// "errors"
)

type user struct {
    email  string
    age int
}

type SuccessResponse struct {
    Status string `json:"status"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}


func handler(w http.ResponseWriter, r *http.Request) {
	// Ambil query param 'name'
    email := r.URL.Query().Get("email")
    if email == "" {
        json.NewEncoder(w).Encode(ErrorResponse{
			Error: "email tidak boleh kosong",
		})
    }

    ageStr := r.URL.Query().Get("age")
    var age int
    var err error

    if ageStr != "" {
		age, err = strconv.Atoi(ageStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "Parameter age harus berupa angka.",
			})
			return
		}

		// Contoh: validasi minimal umur
		if age < 18 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "Minimal umur adalah 18 tahun.",
			})
			return
		}
	}


	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(SuccessResponse{
        Status: "ok",
    })

    // Set header content-type JSON
    // w.Header().Set("Content-Type", "application/json")

    // // Encode ke JSON dan kirim response
    // json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/validate", handler)
	fmt.Println("Server di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}