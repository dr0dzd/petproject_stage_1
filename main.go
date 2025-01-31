package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var allmessages []Message
	geterr := DB.Find(&allmessages)
	if geterr.Error != nil {
		fmt.Fprintf(w, "Data fetch error")
		return
	}
	fmt.Fprintln(w, "All messages:\n", allmessages)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var forproces Message
	decoder := json.NewDecoder(r.Body)
	decerr := decoder.Decode(&forproces)
	if decerr != nil {
		fmt.Fprintf(w, "Error of decoding json")
		return
	}
	createrr := DB.Create(&forproces)
	if createrr.Error != nil {
		fmt.Fprintf(w, "Error add in DB")
		return
	}
	fmt.Fprintf(w, "Message created in Data Base :)")
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	http.ListenAndServe("localhost:8080", router)

}
