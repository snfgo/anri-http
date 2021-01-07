package handlers

import (
	"encoding/json"
	"go-http/models"
	"go-http/storage"
	"io/ioutil"
	"net/http"
)

func CreateUser(service storage.Adder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}

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

		err = service.Add(&user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}
