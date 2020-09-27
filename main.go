package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", categories)
	http.ListenAndServe("0.0.0.0:5050", nil)
}
func categories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCategories())
}
