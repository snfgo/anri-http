package handlers

import (
	"encoding/json"
	"go-http/storage"
	"net/http"
)

func GetUser(service storage.GetterOne) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")

		if len(username) <= 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		user, err := service.GetOne(username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		bytes, err := json.Marshal(user)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(bytes)
	}
}
