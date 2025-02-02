package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		http.Error(w, "Error of decoding json", http.StatusBadRequest)
		return
	}
	if createrr := DB.Create(&forproces); createrr.Error != nil {
		http.Error(w, "Error add in DB", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&forproces); err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, iderr := strconv.Atoi(vars["id"])
	if iderr != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	var forUpdate Message
	if decerr := json.NewDecoder(r.Body).Decode(&forUpdate); decerr != nil {
		http.Error(w, "Error of decoding json", http.StatusBadRequest)
		return
	}

	forUpdate.ID = uint(id)
	if uperr := DB.Model(&Message{}).Where("id = ?", id).Updates(Message{Task: forUpdate.Task, IsDone: forUpdate.IsDone}).Error; uperr != nil {
		http.Error(w, "Update Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if enerr := json.NewEncoder(w).Encode(&forUpdate); enerr != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PUT")
	//router.HandleFunc("api/messages/{id}", PatchMessage).Methods("PATCH")
	//router.HandleFunc("api/messages{id}", DeleteMessage).Methods("DELETE")
	http.ListenAndServe("localhost:8080", router)

}
