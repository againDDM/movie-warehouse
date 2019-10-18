package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println("authMiddleware", r.URL.Path) // to debug
		if user, pass, ok := r.BasicAuth(); !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(Config.AppAdminUser)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(Config.AppAdminPassword)) != 1 {
			log.Println("no auth at", r.URL.Path)
			w.Header().Set("WWW-Authenticate", `Basic realm="`+Config.AppAuthRealm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println("accessLogMiddleware", r.URL.Path) // to debug
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println("panicMiddleware", r.URL.Path) // to debug
		defer func() {
			if err := recover(); err != nil {
				log.Println("RECOVERED", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/films", getFilms).Methods("GET")
	router.HandleFunc("/films/{id}", getFilm).Methods("GET")
	router.HandleFunc("/films", addFilm).Methods("POST")
	router.HandleFunc("/films/{id}", updateFilm).Methods("PUT")
	router.HandleFunc("/films/{id}", deleteFilm).Methods("DELETE")

	siteHandler := authMiddleware(router)
	siteHandler = accessLogMiddleware(siteHandler)
	siteHandler = panicMiddleware(siteHandler)

	log.Println("Start server")
	serverString := fmt.Sprintf("%v:%v", Config.AppListIP, Config.AppPort)
	log.Fatal(http.ListenAndServe(serverString, siteHandler))
}
