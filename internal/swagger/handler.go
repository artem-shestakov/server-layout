package docs

import (
	handlers "api/internal/handlers"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	opts := middleware.RedocOpts{SpecURL: "/docs/swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	router.Handle("/docs", sh)
	router.Handle("/docs/swagger.yml", http.FileServer(http.Dir("./")))
}
