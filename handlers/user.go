package handlers

import (
	"github.com/andhikasamudra/auth-rest-fiber/dto"
	"github.com/andhikasamudra/auth-rest-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	var response dto.JSONResponse
	userModel := models.NewUserModel()
	users, err := userModel.GetUsers(c.Context(), h.DB)
	if err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = dto.BuildGetUserResponse(users)
	response.Message = dto.SuccessMessage
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var request dto.CreateUserRequest
	var response dto.JSONResponse

	err := c.BodyParser(&request)
	if err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	if err := request.Validate(); err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newUser := request.BuildCreateUserRequest()
	err = newUser.CreateUser(c.Context(), h.DB)
	if err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Message = dto.SuccessMessage
	return c.Status(fiber.StatusCreated).JSON(response)
}
