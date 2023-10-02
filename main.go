package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie represents a movie record
type Movie struct {
	ID       string    `json:"id"`
	Imdb     string    `json:"Imdb"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director represents the director of a movie
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // Slice to store movie data

// Handler function to retrieve all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Handler function to delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Find and remove the movie with the matching ID
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Handler function to retrieve a movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Find the movie with the matching ID and send it as the JSON response
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Handler function to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Assign a random ID to the movie and append it to the movies slice
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Handler function to update an existing movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Find the movie with the matching ID, remove it, and replace it with the updated movie
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(10000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter() // Create a new router

	// Add some sample movie data
	movies = append(movies, Movie{ID: "1", Imdb: "9.3", Title: "Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}})
	movies = append(movies, Movie{ID: "2", Imdb: "8.0", Title: "Persuit of Happyness", Director: &Director{Firstname: "Gabriele", Lastname: "Muccino"}})

	// Register the handler functions with their corresponding endpoints
	r.HandleFunc("/", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start the server and log any errors
}
