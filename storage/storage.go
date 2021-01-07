package storage

import (
	"errors"
	"go-http/models"
	"sync"
)

type Adder interface {
	Add(user *models.User) error
}

type Updater interface {
	UpdateOne(username string, user *models.User) (*models.User, error)
}

type Getter interface {
	GetAll() []*models.User
}

type GetterOne interface {
	GetOne(username string) (*models.User, error)
}

type Deleter interface {
	DeleteUser(username string) error
}

type Users struct {
	mx      sync.RWMutex
	storage map[string]*models.User
}

func New() *Users {
	return &Users{
		storage: make(map[string]*models.User),
	}
}

func (users *Users) Add(user *models.User) error {
	users.mx.Lock()
	defer users.mx.Unlock()
	if _, ok := users.storage[user.Username]; ok {
		return errors.New("user already exists")
	}

	users.storage[user.Username] = user

	return nil
}

func (users *Users) GetAll() []*models.User {
	users.mx.RLock()
	defer users.mx.RUnlock()
	list := make([]*models.User, 0, len(users.storage))

	for _, value := range users.storage {
		list = append(list, value)
	}

	return list
}

func (users *Users) GetOne(username string) (*models.User, error) {
	users.mx.RLock()
	defer users.mx.RUnlock()
	if user, ok := users.storage[username]; ok {
		return user, nil
	}

	return nil, errors.New("user not exists")
}

func (users *Users) UpdateOne(username string, user *models.User) (*models.User, error) {
	users.mx.Lock()
	defer users.mx.Unlock()
	if _, ok := users.storage[username]; !ok {
		return nil, errors.New("user not exists")
	}

	delete(users.storage, username)

	users.storage[user.Username] = user

	return user, nil
}

func (users *Users) DeleteUser(username string) error {
	users.mx.Lock()
	defer users.mx.Unlock()

	if _, ok := users.storage[username]; !ok {
		return errors.New("user not exists")
	}

	delete(users.storage, username)

	return nil
}
