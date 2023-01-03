package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Msaorc/SearchZipCode/server/packages"
)

func main() {
	http.HandleFunc("/search", Search)
	http.ListenAndServe(":9000", nil)
}

func Search(w http.ResponseWriter, r *http.Request) {
	zipCode := r.URL.Query().Get("zipCode")
	result, error := packages.SearchZipCode(zipCode)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(error)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
