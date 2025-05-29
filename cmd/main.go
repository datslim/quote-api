package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type UserData struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

var UserDatas []UserData

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", postQuote).Methods("POST")
	router.HandleFunc("/quotes", getAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", getRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", deleteQuote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	if author == "" {
		json.NewEncoder(w).Encode(UserDatas)
	}

	var filtredQuotes []UserData
	for _, quote := range UserDatas {
		if quote.Author == author {
			filtredQuotes = append(filtredQuotes, quote)
		}
	}

	json.NewEncoder(w).Encode(filtredQuotes)
}

func postQuote(w http.ResponseWriter, r *http.Request) {
	var newData UserData
	var currentID int = 0
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(UserDatas) == 0 {
		currentID = 1
	} else {
		currentID = len(UserDatas) + 1
	}
	newData.ID = currentID
	UserDatas = append(UserDatas, newData)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newData)
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	if len(UserDatas) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}
	randomNumber := rand.Intn(len(UserDatas))
	json.NewEncoder(w).Encode(UserDatas[randomNumber])
}

func deleteQuote(w http.ResponseWriter, r *http.Request) {

}
