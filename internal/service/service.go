package service

import (
	"ablufus/internal/entities"
	"ablufus/internal/repository"
	"fmt"
)

type Service interface {
	Post(u entities.UserRequest) (*entities.UserResponse, error)
	List(ids []string, limit int, page int) ([]*entities.UserResponse, error)
	Update(id string, amount float64) error
}

type UserService struct {
	r repository.Repository
}

func New(r repository.Repository) Service {
	return &UserService{r}
}

func (s *UserService) Post(u entities.UserRequest) (*entities.UserResponse, error) {
	usr, err := entities.ToUser(u)
	if err != nil {
		return nil, err
	}
	res, err := s.r.Post(usr)
	if err != nil {
		return nil, err
	}
	return entities.ToUserResponse(res), nil
}

func (s *UserService) List(ids []string, limit, page int) ([]*entities.UserResponse, error) {
	res, err := s.r.List(ids, limit, page)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response := []*entities.UserResponse{}
	for _, rs := range res {
		response = append(response, entities.ToUserResponse(rs))
	}
	return response, nil
}

func (s *UserService) Update(id string, amount float64) error {
	return s.r.Update(id, amount)
}
