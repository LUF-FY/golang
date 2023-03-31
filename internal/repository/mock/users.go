package mock

import (
	"errors"
	"golang/internal/models"
	"time"
)

type UserRepo struct {
	lastid  int64
	storage map[interface{}]*models.User
}

func NewRepo() *UserRepo {
	return &UserRepo{
		lastid:  0,
		storage: make(map[interface{}]*models.User),
	}
}

func (r *UserRepo) Create(name, email, password string) (int64, error) {
	user := &models.User{
		Id:        r.lastid,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	r.lastid += 1
	r.storage[name] = user
	r.storage[email] = user
	r.storage[password] = user

	return user.Id, nil
}

func (r *UserRepo) Get(name, password string) (*models.User, error) {
	user, ok := r.storage[password]
	if !ok || user.Name != name {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepo) Update(userId int64, field, value string) error {
	user, ok := r.storage[userId]
	if !ok {
		return errors.New("User not found")
	}

	switch field {
	case "name":
		user.Name = value
	case "email":
		user.Email = value
	default:
		return errors.New("invalid field")
	}

	return nil
}

func (r *UserRepo) Delete(userId int64) error {
	if _, ok := r.storage[userId]; !ok {
		return errors.New("User not found")
	}

	delete(r.storage, userId)
	return nil
}
