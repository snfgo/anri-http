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

const getUserTestUrl = "/users/get"

func createTestGetOneUserServer(storage *storage.Users) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc(getUserTestUrl, GetUser(storage))
	return router
}

func TestGetOneUser(t *testing.T) {
	storage := storage.New()
	srv := httptest.NewServer(createTestGetOneUserServer(storage))
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
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", srv.URL, getUserTestUrl), bytes.NewReader(nil))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusBadRequest)

	})

	t.Run("get non-existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", srv.URL, getUserTestUrl), bytes.NewReader(nil))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusBadRequest)

	})

	t.Run("get existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?username=%s", srv.URL, getUserTestUrl, user.Username), bytes.NewReader(nil))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusOK)

	})
}
