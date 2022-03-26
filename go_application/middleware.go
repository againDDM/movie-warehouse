package main

import (
	"crypto/subtle"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
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
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", ping)
	r.Get("/films", getFilms)
	r.Get("/films/{id}", getFilm)
	r.Post("/films", addFilm)
	r.Put("/films/{id}", updateFilm)
	r.Delete("/films/{id}", deleteFilm)

	log.Println("Start server")
	serverString := fmt.Sprintf("%v:%v", Config.AppListIP, Config.AppPort)
	log.Fatal(http.ListenAndServe(serverString, panicMiddleware(accessLogMiddleware(authMiddleware(r)))))
}
