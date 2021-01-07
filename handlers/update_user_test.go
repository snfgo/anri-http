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

const updateUserTestUrl = "/users/update"

func createTestUpdateUserServer(storage *storage.Users) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc(updateUserTestUrl, UpdateUser(storage))
	return router
}

func TestUpdateUser(t *testing.T) {
	storage := storage.New()
	srv := httptest.NewServer(createTestUpdateUserServer(storage))
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

	t.Run("query without username", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s?", srv.URL, updateUserTestUrl), bytes.NewReader(data))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusBadRequest)

	})

	t.Run("update non-existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s?username=%s", srv.URL, updateUserTestUrl, "someuser"), bytes.NewReader(data))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusNotFound)

	})

	t.Run("update existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s?username=%s", srv.URL, updateUserTestUrl, user.Username), bytes.NewReader(data))
		_ = storage.Add(&user)

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusOK)
	})
}
