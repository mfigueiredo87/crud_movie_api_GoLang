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

// criando struct por nao usar um banco de dados
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// slice para armazenar os filmes
var movies []Movie

// function to get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get movie by id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// function to create Movie

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Dados inválidos"})
		return
	}

	// Garante que o ID é único
	movie.ID = strconv.Itoa(rand.Intn(1000000000))

	// Inicializa director se não vier no JSON
	if movie.Director == nil {
		movie.Director = &Director{
			Firstname: "",
			Lastname:  "",
		}
	}

	// Inicializa os campos vazios se necessário
	if movie.Isbn == "" {
		movie.Isbn = "N/A"
	}
	if movie.Title == "" {
		movie.Title = "Sem Título"
	}

	// Adiciona o movie ao slice
	movies = append(movies, movie)

	// Retorna mensagem de sucesso com todos os campos
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie adicionado com sucesso",
		"movie":   movie,
	})
}


// funcao para atualizar o movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// remover movie antigo
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// funcao para apagar o movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie apagado com sucesso",
		"moviews": movies,
	})
}

// funcao principal
func main() {
	r := mux.NewRouter()

	// adicionando alguns filmes ao slice
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// definindo as funcoes, endpoints e respectivos metodos
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// iniciando o servidor
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
