package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func service() http.Handler {
	r := chi.NewRouter()
	// TODO: Use zap.logger https://pkg.go.dev/go.uber.org/zap example with logrus https://github.com/go-chi/chi/blob/master/_examples/logging/main.go
	r.Use(middleware.CleanPath, middleware.GetHead, middleware.RequestID, middleware.RealIP)

	r.With(middleware.Logger, middleware.Recoverer).Group(func(r chi.Router) {
		r.Get("/ping", ping)
	})

	r.With(middleware.BasicAuth(Config.AppAuthRealm, map[string]string{Config.AppAdminUser: Config.AppAdminPassword}),
		middleware.Logger, middleware.Recoverer).Group(func(r chi.Router) {

		r.Get("/films", getFilms)
		r.Get("/films/{id}", getFilm)
		r.Post("/films", addFilm)
		r.Put("/films/{id}", updateFilm)
		r.Delete("/films/{id}", deleteFilm)
	})

	return r
}

func Run() {
	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", Config.AppListIP, Config.AppPort),
		Handler: service(),
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func Ping() error {
	client := http.Client{Timeout: time.Second}
	response, err := client.Get(fmt.Sprintf("http://localhost:%v/ping", Config.AppPort))
	if err != nil {
		return err
	}
	if response != nil {
		defer func() {
			if err := response.Body.Close(); err != nil {
				log.Printf("fail close body: %s", err)
			}
		}()
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("expect status %v, got %v", http.StatusOK, response.StatusCode)
	}
	log.Println("It`s alive")
	return nil
}
