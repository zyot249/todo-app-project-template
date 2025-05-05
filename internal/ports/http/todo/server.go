package todo

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RunServer(createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddleware(apiRouter)
	setSwaggerDoc(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", createHandler(apiRouter))

	logrus.Info("Starting server on port 8080")
	logrus.Info("API Documentation: ", "http://localhost:8080/doc/index.html")

	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		logrus.WithError(err).Panic("Failed to start server")
	}
}

func setSwaggerDoc(router *chi.Mux) {
	router.Get("/todo/doc.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/ports/http/todo/openapi.yml")
	})
	router.Get("/todo/doc/*", httpSwagger.Handler(
		httpSwagger.URL("/todo/doc.yml"),
	))
}

func setMiddleware(router *chi.Mux) {
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	addCorsMiddleware(router)

	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.SetHeader("X-Content-Type-Options", "nosniff"))
	router.Use(middleware.SetHeader("X-Frame-Options", "DENY"))

	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(cors.Handler)
}
