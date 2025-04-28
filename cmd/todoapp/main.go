package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	apiRouter := chi.NewRouter()
	apiRouter.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	router.Mount("/api", apiRouter)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
