package dto

import (
	"errors"
	"github.com/andhikasamudra/auth-rest-fiber/models"
	"github.com/andhikasamudra/auth-rest-fiber/util"
	"github.com/google/uuid"
)

type GetUserResponse struct {
	Email string    `json:"email"`
	Name  string    `json:"name"`
	GUID  uuid.UUID `json:"guid"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func BuildGetUserResponse(user []models.User) []GetUserResponse {
	var response []GetUserResponse

	for _, value := range user {
		response = append(response, GetUserResponse{
			Email: value.Email,
			Name:  value.Name,
			GUID:  value.GUID,
		})
	}

	return response
}

func (r CreateUserRequest) BuildCreateUserRequest() models.User {
	return models.User{
		Email:    r.Email,
		Name:     r.Name,
		Password: util.Encode([]byte(r.Password)),
	}
}

func (r CreateUserRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}

	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
