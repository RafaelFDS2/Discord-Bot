package entities

import (
	"ablufus/exceptions"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	Amount    float64
}

type UserResponse struct {
	ID        string    `json:"userId"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UserRequest struct {
	ID string `json:"userId" validate:"required"`
}

type UserPatchRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}

func ToUser(u UserRequest) (*User, error) {
	if err := Validate(u); err != nil {
		return nil, err
	}
	return &User{ID: u.ID, Amount: 0}, nil
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Amount:    u.Amount,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func Validate(u interface{}) error {
	validate := validator.New()
	res := validate.Struct(u)
	if res != nil {
		for _, err := range res.(validator.ValidationErrors) {
			exceptions.ErrValidation = NewException(err)
			break
		}
		return exceptions.New(exceptions.ErrValidation, res)
	}
	return nil
}

func NewException(err validator.FieldError) error {
	var message string
	switch err.ActualTag() {
	case "required":
		message = fmt.Sprintf("user: missing %s", err.Field())
	case "min":
		message = fmt.Sprintf("user: %s doesn't match to the min size", err.Field())
	case "max":
		message = fmt.Sprintf("user: %s doesn't fit in max size", err.Field())
	}
	return errors.New(message)
}
