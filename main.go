package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/categories", listCategories).Methods("GET")
	r.HandleFunc("/", listWords).Methods("GET")
	r.HandleFunc("/", addWord).Methods("POST")
	r.HandleFunc("/", editWord).Methods("PUT")
	r.HandleFunc("/", eraseWord).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5050", r))
}
func listWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getWords())
}
func listCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCategories())
}
func addWord(w http.ResponseWriter, r *http.Request) {
	var word Word
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&word)
	createWord(word)
}

func editWord(w http.ResponseWriter, r *http.Request) {
	var word Word
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&word)
	updateWord(word)
}
func eraseWord(w http.ResponseWriter, r *http.Request) {
	var word Word
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&word)
	deleteWord(word)
}
