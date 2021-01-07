package helpers

import (
	"net/http"
)

const CorsMethods = "POST, GET, OPTIONS, PUT, DELETE"
const CorsOrigin = "*"
const CorsHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization"

func WithMethod(fn func(http.ResponseWriter, *http.Request), method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		fn(w, r)
	}
}

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", CorsOrigin)
		w.Header().Set("Access-Control-Allow-Methods", CorsMethods)
		w.Header().Set("Access-Control-Allow-Headers", CorsHeaders)
		h.ServeHTTP(w, r)
	})
}
