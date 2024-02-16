package repository

import (
	"ablufus/exceptions"
	"ablufus/internal/database"
	"ablufus/internal/entities"
	"fmt"
)

type Repository interface {
	Post(u *entities.User) (*entities.User, error)
	List(ids []string, limit int, page int) ([]*entities.User, error)
	Update(id string, amount float64) error
}

type UserRepository struct {
	db database.Database
}

func New(db database.Database) Repository {
	return &UserRepository{db}
}

func (r *UserRepository) Post(u *entities.User) (*entities.User, error) {
	res, err := r.db.Post(u)
	if err != nil {
		fmt.Println(err.Error())
		return nil, exceptions.New(exceptions.ErrUserAlreadyExists, err)
	}
	return res, nil
}

func (r *UserRepository) List(ids []string, limit int, page int) ([]*entities.User, error) {
	res, err := r.db.List(ids, limit, page)
	if err != nil {
		return nil, exceptions.New(exceptions.ErrBadData, err) // Update exceptions to handle not found
	}
	return res, nil
}

func (r *UserRepository) Update(id string, amount float64) error {
	if err := r.db.Update(id, amount); err != nil {
		return exceptions.New(exceptions.ErrBadData, err) // Update exceptions to handle not found
	}
	return nil
}
