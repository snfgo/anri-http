package handlers

import (
	"bytes"
	"fmt"
	"go-http/models"
	"go-http/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const deleteUserTestUrl = "/users/delete"

func createDeleteTestUserServer(storage *storage.Users) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc(deleteUserTestUrl, DeleteUser(storage))
	return router
}

func TestDeleteUser(t *testing.T) {
	storage := storage.New()

	srv := httptest.NewServer(createDeleteTestUserServer(storage))
	assert := assert.New(t)
	defer srv.Close()

	user := models.User{
		Username: "John",
		Age:      24,
		Email:    "john@test.com",
		Gender:   "male",
	}

	_ = storage.Add(&user)

	t.Run("query without username", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", srv.URL, deleteUserTestUrl), bytes.NewReader(nil))

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusBadRequest)
	})

	t.Run("delete existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s?username=%s", srv.URL, deleteUserTestUrl, user.Username), bytes.NewReader(nil))

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusNoContent)
	})

	t.Run("delete non-existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s?username=%s", srv.URL, deleteUserTestUrl, user.Username), bytes.NewReader(nil))

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusNotFound)
	})
}
