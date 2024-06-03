package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"newproj/handlers"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Get("/foo", handlers.Make(handlers.HandleFoo))

	listendAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "addr", listendAddr)

	http.ListenAndServe(listendAddr, router)

}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo")
}
