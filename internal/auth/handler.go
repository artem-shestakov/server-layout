package auth

import (
	apperror "api/internal/apperrors"
	handlers "api/internal/handlers"
	"encoding/json"
	"net/http"

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
	authRouter := router.PathPrefix(("/auth")).Subrouter()
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
	// 		 500: errorResponse
	// vdcRouter.HandleFunc("/uidbysn", apperror.Middleware(h.GetBySn)).Methods(http.MethodGet)
	authRouter.HandleFunc("/signup", apperror.Middleware(h.SignUp)).Methods(http.MethodPost)
	authRouter.HandleFunc("/signin", apperror.Middleware(h.SignIn)).Methods(http.MethodPost)
}

func (h *handler) SignUp(rw http.ResponseWriter, r *http.Request) error {
	var user CreateUser
	// Get data from request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return apperror.NewError(err, http.StatusBadRequest, "Incorrect request params", err.Error())
	}
	// Create user in database
	newUser, err := h.service.CreateUser(&user)
	if err != nil {
		return err
	}
	// Response
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		return apperror.NewError(err, http.StatusInternalServerError, "Can't marshal of new user data", err.Error())
	}
	rw.Write(jsonBytes)
	return nil
}

func (h *handler) SignIn(rw http.ResponseWriter, r *http.Request) error {
	var signInInput SignInInput
	err := json.NewDecoder(r.Body).Decode(&signInInput)
	if err != nil {
		return apperror.NewError(err, http.StatusBadRequest, "Incorrect request params", err.Error())
	}
	JWTToken, err := h.service.GenerateJWT(signInInput.Email, signInInput.Password)
	if err != nil {
		return err
	}

	token := JWTResponse{
		Token: JWTToken,
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	jsonBytes, err := json.Marshal(token)
	if err != nil {
		return apperror.NewError(err, http.StatusInternalServerError, "Can't marshal of new user data", err.Error())
	}
	rw.Write(jsonBytes)
	return nil
}
