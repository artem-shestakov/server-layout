package apperror

import (
	"api/internal/prom"
	"errors"
	"net/http"
	"time"
)

type handlerFunc func(rw http.ResponseWriter, r *http.Request) error

// Middleware error handler
func Middleware(h handlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var appErr *Error
		begin := time.Now()
		err := h(rw, r)
		if err != nil {
			if errors.As(err, &appErr) {
				err := err.(*Error)
				err.Send(rw)
				prom.UpdateMetrics(err.StatusCode, r, begin)
				return
			}
			teaPotError(err).Send(rw)
			prom.UpdateMetrics(http.StatusTeapot, r, begin)
			return
		}
		prom.UpdateMetrics(200, r, begin)
	}
}
