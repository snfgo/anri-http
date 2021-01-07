package main

import (
	"fmt"
	"go-http/handlers"
	"go-http/helpers"
	"go-http/storage"
	"net/http"
)

const port = "4040"

func main() {
	srv := http.NewServeMux()
	storage := storage.New()
	srv.HandleFunc("/users/create", helpers.WithMethod(handlers.CreateUser(storage), http.MethodPost))
	srv.HandleFunc("/users/list", helpers.WithMethod(handlers.GetUsers(storage), http.MethodGet))
	srv.HandleFunc("/users/update", helpers.WithMethod(handlers.UpdateUser(storage), http.MethodPut))
	srv.HandleFunc("/users/delete", helpers.WithMethod(handlers.DeleteUser(storage), http.MethodDelete))
	srv.HandleFunc("/users/get", helpers.WithMethod(handlers.GetUser(storage), http.MethodGet))

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), helpers.Cors(srv))
}
