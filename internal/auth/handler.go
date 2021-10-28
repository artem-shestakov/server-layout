package auth

import (
	handlers "api/internal/handlers"

	"github.com/gorilla/mux"
)

type handler struct {
	service AuthService
}

func NewHandler(service AuthService) handlers.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(router *mux.Router) {
	// vdcRouter := router.PathPrefix("/vdc").Subrouter()
	// swagger:route GET /vdc/uidbysn vdc uidbysn
	// Returns information about camera by its serial number
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: cameraResponse
	// 		 400: errorResponse
	// 		 404: errorResponse
	// 		 418: errorResponse
	// 		 422: errorResponse
	// 		 500: errorResponse
	// vdcRouter.HandleFunc("/uidbysn", apperror.Middleware(h.GetBySn)).Methods(http.MethodGet)
}
