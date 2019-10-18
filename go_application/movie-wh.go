package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Film is a data structure for marshaling convenience.
type Film struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Director is a data structure for marshaling convenience.
type Director struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func getFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	films := []*Film{}
	rows, err := db.Query("SELECT id, name, description FROM films")
	defer rows.Close()
	if err != nil {
		log.Printf("Fail to get films from database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	for rows.Next() {
		movie := &Film{}
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Description)
		if err != nil {
			log.Printf("Fail to get films from database :: %v /n", err)
			http.Error(w, "Internal server error", 500)
			return
		}
		films = append(films, movie)
	}
	json.NewEncoder(w).Encode(films)
}

func getFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	targetID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Invalid ID")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	row := db.QueryRow("SELECT id, name, description FROM films WHERE id=$1", targetID)
	if err != nil {
		log.Printf("Fail to get film from database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	movie := &Film{}
	switch err := row.Scan(&movie.ID, &movie.Name, &movie.Description); err {
	case nil:
		json.NewEncoder(w).Encode(movie)
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
	default:
		log.Printf("Fail to get film from database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	}
}

func addFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var film Film
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		log.Println("Can`t decode user`s request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = db.Exec("INSERT INTO films (name, description) VALUES($1, $2)",
		film.Name, film.Description)
	if err != nil {
		log.Printf("Fail to add film to database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func deleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	targetID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Invalid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	aff, err := db.Exec("DELETE FROM films WHERE id=$1", targetID)
	switch {
	case err != nil:
		log.Printf("Fail to delete film from database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	case aff == nil:
		log.Println("Nothing to delete")
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusNoContent)
	}
}

func updateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	targetID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Invalid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var film Film
	err = json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		log.Println("Can`t decode user`s request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	aff, err := db.Exec("UPDATE films  SET name=$1, description=$2 WHERE id=$3",
		film.Name, film.Description, targetID)
	switch {
	case err != nil:
		log.Printf("Fail to update film in database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	case aff == nil:
		log.Println("Nothing to update")
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusNoContent)
	}
}
