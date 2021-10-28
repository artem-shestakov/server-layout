package apperror

import "net/http"

func MethodNotAllowed() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		methodNotAllowed := NewError(nil, 405, "Method not allowed", "")
		methodNotAllowed.Send(rw)
	})
}

func NotFound() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		methodNotAllowed := NewError(nil, 404, "Not found. Check URL", "")
		methodNotAllowed.Send(rw)
	})
}
