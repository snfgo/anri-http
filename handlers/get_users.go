package handlers

import (
	"encoding/json"
	"go-http/storage"
	"net/http"
)

func GetUsers(service storage.Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		users := service.GetAll()

		bytes, err := json.Marshal(users)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write(bytes)
	}
}
