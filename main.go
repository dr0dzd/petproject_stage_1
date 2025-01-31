package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var allmessages []Message
	geterr := DB.Find(&allmessages)
	if geterr.Error != nil {
		http.Error(w, "Data fetch error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if senderr := encoder.Encode(allmessages); senderr != nil {
		http.Error(w, "Error of transformation", http.StatusInternalServerError)
		return
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var forproces Message
	decoder := json.NewDecoder(r.Body)
	if decerr := decoder.Decode(&forproces); decerr != nil {
		http.Error(w, "Error of decoding json", http.StatusInternalServerError)
		return
	}
	if createrr := DB.Create(&forproces); createrr.Error != nil {
		http.Error(w, "Error add in DB", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&forproces); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	http.ListenAndServe("localhost:8080", router)

}
