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
	if geterr != nil {
		fmt.Errorf("Ошибка извлечения пользователей: %w", geterr)
		return
	}
	fmt.Fprintln(w, "All messages:", allmessages)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var forproces Message
	decoder := json.NewDecoder(r.Body)
	decerr := decoder.Decode(&forproces)
	if decerr != nil {
		fmt.Errorf("Ошибка: %w", decerr)
		return
	}
	createrr := DB.Create(&forproces)
	if createrr != nil {
		fmt.Errorf("Ошибка добавления в БД: %w", createrr)
		return
	}
	fmt.Fprintln(w, "Message created in Data Base :)")
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	http.ListenAndServe("localhost:8080", router)

}
