package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Jsonheader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//Allow origin dari semua sumber
		// Hapus jika sudah tidak digunakan
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next(w, r)
	}
}

func Jsoncontent(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Jsoncontenterror(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Jsoncontent(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	w.WriteHeader(statusCode)
}
