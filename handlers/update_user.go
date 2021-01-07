package handlers

import (
	"encoding/json"
	"go-http/models"
	"go-http/storage"
	"io/ioutil"
	"net/http"
)

func UpdateUser(service storage.Updater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

		username := r.URL.Query().Get("username")

		if len(username) <= 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(data, &user)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		res, err := service.UpdateOne(username, &user)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		bytes, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(bytes)
		}
	}
}
