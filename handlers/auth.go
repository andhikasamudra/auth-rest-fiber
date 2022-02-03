package handlers

import (
	"github.com/andhikasamudra/auth-rest-fiber/dto"
	"github.com/andhikasamudra/auth-rest-fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest
	var response dto.JSONResponse
	var userModel models.User

	if err := c.BodyParser(&request); err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := request.Validate(); err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	currentUser, err := userModel.GetUserByEmail(c.Context(), h.DB, request.Email)
	if err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(request.Password)) != nil {
		response.Message = "Invalid email or password"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	jwtClaims := dto.JWTClaims{
		GUID: currentUser.GUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	signedToken, err := token.SignedString([]byte(h.Config.TokenSecret()))
	if err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = signedToken
	return c.JSON(response)
}
