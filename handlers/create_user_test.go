package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-http/models"
	"go-http/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const createUserTestUrl = "/users/create"

func createTestAddUserServer() http.Handler {
	storage := storage.New()
	router := http.NewServeMux()
	router.HandleFunc(createUserTestUrl, CreateUser(storage))
	return router
}

func TestCreateUser(t *testing.T) {
	srv := httptest.NewServer(createTestAddUserServer())
	assert := assert.New(t)
	defer srv.Close()

	user := models.User{
		Username: "John",
		Age:      24,
		Email:    "john@test.com",
		Gender:   "male",
	}

	data, err := json.Marshal(user)

	if err != nil {
		t.Error(err)
	}

	t.Run("create new user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", srv.URL, createUserTestUrl), bytes.NewReader(data))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusCreated)

	})

	t.Run("create existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", srv.URL, createUserTestUrl), bytes.NewReader(data))

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusBadRequest)
	})
}
