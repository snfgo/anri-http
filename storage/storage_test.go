package storage

import (
	"go-http/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersStorage(t *testing.T) {
	users := New()
	assert := assert.New(t)

	user := models.User{
		Username: "John",
		Age:      24,
		Email:    "john@test.com",
		Gender:   "male",
	}

	newUser := models.User{
		Username: "Jane",
		Age:      24,
		Email:    "jane@test.com",
		Gender:   "female",
	}

	t.Run("add user", func(t *testing.T) {
		err := users.Add(&user)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(len(users.storage), 1)
	})

	t.Run("get users", func(t *testing.T) {
		list := users.GetAll()
		assert.Equal(len(list), 1)
	})

	t.Run("get user that does not exist", func(t *testing.T) {
		_, err := users.GetOne("some user")
		assert.NotNil(err)
	})

	t.Run("get user that exists", func(t *testing.T) {
		res, err := users.GetOne(user.Username)

		assert.Nil(err)
		assert.Equal(res.Username, user.Username)
	})

	t.Run("update user that does not exist", func(t *testing.T) {
		_, err := users.UpdateOne("some user", &newUser)
		assert.NotNil(err)
	})

	t.Run("update user that does not exist", func(t *testing.T) {

		res, err := users.UpdateOne(user.Username, &newUser)

		assert.Nil(err)
		assert.Equal(res.Username, newUser.Username)
	})

	t.Run("delete user that does not exist", func(t *testing.T) {
		err := users.DeleteUser("some user")
		assert.NotNil(err)
	})

	t.Run("delete user that exists", func(t *testing.T) {
		err := users.DeleteUser(newUser.Username)
		assert.Nil(err)
	})
}
