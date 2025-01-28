package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if task == "" {
		fmt.Fprintln(w, "Hello, world!")
	} else {
		fmt.Fprintf(w, "Hello, %s\n", task)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		//fmt.Fprintln(w, "Ошибка парсинга (я очень крутой узнал новое слово)")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task = reqBody.Message

	fmt.Fprintf(w, "Task set to: %s", task)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", GetHandler).Methods("GET")

	router.HandleFunc("/api/task", PostHandler).Methods("POST")

	http.ListenAndServe("localhost:8080", router)

}
