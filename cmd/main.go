package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type UserData struct {
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
	router.HandleFunc("/quotes", deleteQuote).Methods("DELETE")
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
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserDatas = append(UserDatas, newData)
	w.WriteHeader(http.StatusCreated)
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
