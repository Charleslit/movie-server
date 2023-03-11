package main

import (
	"encoding/json"
	"fmt"
	"log" //for logging errors
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	title    string    `json:"title"`
	director *director `json:"Director"`
}
type director struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var movies []movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
}
func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var Movie movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.ID = strconv.Itoa((rand.Intn(1000000)))
	movies = append(movies, Movie)
	json.NewEncoder(w).Encode(Movie)

}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			var Movie movie
			_ = json.NewDecoder(r.Body).Decode(&Movie)
			Movie.ID = params["id"]
			movies = append(movies, Movie)
			json.NewEncoder(w).Encode(Movie)
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", isbn: "456756", title: "movie one", director: &director{firstname: "john", lastname: "ogwoka "}})
	movies = append(movies, movie{ID: "2", isbn: "4567898", title: "movietwoo", director: &director{firstname: "charles", lastname: "duoe"}})
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{ID} ik", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deletemovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}
