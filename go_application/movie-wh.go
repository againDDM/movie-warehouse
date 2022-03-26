package main

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
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
	rows, err := db.Query("SELECT id, name, description FROM films ORDER BY id")
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
	strID := chi.URLParam(r, "id")
	targetID, err := strconv.Atoi(strID)
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
	if film.Name == "" && film.Description == "" {
		log.Println("Empty insert")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := db.Exec("INSERT INTO films (name, description) VALUES($1, $2)",
		film.Name, film.Description)
	if err != nil {
		log.Printf("Fail to add film to database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Query error :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	} else {
		log.Printf("Added %v film(s) to database/n", rowsAffected)
		w.WriteHeader(http.StatusCreated)
	}
}

func deleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	strID := chi.URLParam(r, "id")
	targetID, err := strconv.Atoi(strID)
	if err != nil {
		log.Println("Invalid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := db.Exec("DELETE FROM films WHERE id=$1", targetID)
	if err != nil {
		log.Printf("Fail to delete film from database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Query error :: %v /n", err)
		http.Error(w, "Internal server error", 500)
	} else if rowsAffected == 0 {
		log.Println("Nothing to delete from database")
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Printf("Deleted %v film(s) from database/n", rowsAffected)
		w.WriteHeader(http.StatusNoContent)
	}
}

func updateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	strID := chi.URLParam(r, "id")
	targetID, err := strconv.Atoi(strID)
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
	var result sql.Result
	switch {
	case film.Name == "" && film.Description == "":
		log.Println("Empty update")
		w.WriteHeader(http.StatusBadRequest)
		return
	case film.Name == "":
		result, err = db.Exec("UPDATE films  SET description=$1 WHERE id=$2",
			film.Description, targetID)
	case film.Description == "":
		result, err = db.Exec("UPDATE films  SET name=$1 WHERE id=$2",
			film.Name, targetID)
	default:
		result, err = db.Exec("UPDATE films  SET name=$1, description=$2 WHERE id=$3",
			film.Name, film.Description, targetID)
	}
	if err != nil {
		log.Printf("Fail to update film in database :: %v /n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Query error :: %v /n", err)
		http.Error(w, "Internal server error", 500)
		return
	} else if rowsAffected == 0 {
		log.Println("Nothing to update in database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		log.Printf("Updated %v film(s) from database/n", rowsAffected)
		w.WriteHeader(http.StatusAccepted)
	}
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("pong\n"))
	w.WriteHeader(http.StatusOK)
}
