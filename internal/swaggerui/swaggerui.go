package swaggerui

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SwaggerUIServer struct {
}

func NewSwaggerUIServer() SwaggerUIServer {
	return SwaggerUIServer{}
}

func (s SwaggerUIServer) GetPort() string {
	return "8000"
}

func (s SwaggerUIServer) GetServiceName() string {
	return "swaggerui"
}

func (s SwaggerUIServer) RegisterMiddlewares(router *chi.Mux) {
}

func (s SwaggerUIServer) CreateApiHandler(router *chi.Mux) http.Handler {
	return router
}

func (s SwaggerUIServer) CreateRootHandler(router *chi.Mux) http.Handler {
	router.Group(func(router chi.Router) {
		router.Get("/config.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "third_party/swaggerui/config.json")
		})
	})

	path := "/doc"
	root := http.Dir("third_party/swaggerui")
	fs := http.StripPrefix(path, http.FileServer(root))

	router.Group(func(router chi.Router) {
		router.Get(path+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		}))
	})

	return router
}

func (s SwaggerUIServer) HaveApiSwaggerDoc() bool {
	return false
}
