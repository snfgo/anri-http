package handlers

import (
	"go-http/storage"
	"net/http"
)

func DeleteUser(service storage.Deleter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")

		if len(username) <= 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err := service.DeleteUser(username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
