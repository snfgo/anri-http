package handlers

import (
	"bytes"
	"fmt"
	"go-http/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const getUsersTestUrl = "/users/list"

func createTestGetUsersServer() http.Handler {
	storage := storage.New()
	router := http.NewServeMux()
	router.HandleFunc(getUsersTestUrl, GetUsers(storage))
	return router
}

func TestGetUsers(t *testing.T) {
	srv := httptest.NewServer(createTestGetUsersServer())
	assert := assert.New(t)
	defer srv.Close()

	t.Run("get all users", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", srv.URL, getUsersTestUrl), bytes.NewReader(nil))
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(res.StatusCode, http.StatusOK)
	})
}
