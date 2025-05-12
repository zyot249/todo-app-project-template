package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"todo-app/internal/common/core/http/constants"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type ShynobiServer interface {
	GetPort() string
	GetHost() string
	GetServiceName() string
	RegisterMiddlewares(router *chi.Mux)
	CreateApiHandler(router *chi.Mux) http.Handler
	CreateRootHandler(router *chi.Mux) http.Handler
	HaveApiSwaggerDoc() bool
}

func Start(server ShynobiServer) {
	addr := fmt.Sprintf("%s:%s", server.GetHost(), server.GetPort())
	endpoint := fmt.Sprintf("http://%s", addr)

	apiRouter := chi.NewRouter()

	registerDefaultMiddlewares(apiRouter)
	server.RegisterMiddlewares(apiRouter)

	if server.HaveApiSwaggerDoc() {
		registerSwaggerDoc(apiRouter, server.GetServiceName())
		logrus.Info("API Swagger Doc: ", endpoint+constants.BaseApiPath+fmt.Sprintf(constants.SwaggerEndpointFormat, server.GetServiceName()))
	}

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/", server.CreateRootHandler(rootRouter))
	rootRouter.Mount(constants.BaseApiPath, server.CreateApiHandler(apiRouter))

	logrus.Info("Starting HTTP server on ", endpoint)

	if err := http.ListenAndServe(addr, rootRouter); err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
}

func registerSwaggerDoc(router *chi.Mux, serviceName string) {
	var docPath = fmt.Sprintf(constants.SwaggerDocPathFormat, serviceName)
	var docFilePath = fmt.Sprintf(constants.SwaggerDocFilePathFormat, serviceName)

	router.Get(fmt.Sprintf(constants.DocPathFormat, serviceName), httpSwagger.Handler(httpSwagger.URL(constants.BaseApiPath+docPath)))
	router.Get(docPath, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, docFilePath)
	})
}

func registerDefaultMiddlewares(router *chi.Mux) {
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	addCorsMiddleware(router)

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
