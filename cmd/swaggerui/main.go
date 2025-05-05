package main

import (
	"net/http"
	"os"
	"todo-app/pkg/core/logs"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	logs.Init()
	logger := logrus.NewEntry(logrus.StandardLogger())

	baseDir, _ := os.Getwd()
	logger.Info("Base directory: ", baseDir)

	root := chi.NewRouter()

	// Serve swagger-config.json
	root.Get("/config.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "third_party/swaggerui/config.json")
	})

	// Serve Swagger UI static files
	FileServer(root, "/api/doc/", http.Dir("third_party/swaggerui"))

	http.ListenAndServe(":8000", root)
}

// Helper to serve static files
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if path[len(path)-1] != '/' {
		panic("path must end with /")
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
