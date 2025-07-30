package controllers

import (
	"first_api_golang/models"
	"first_api_golang/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func (authcontroller *AuthController) Register(c *fiber.Ctx) error {
	var request models.RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Data user required")
	}
	var existingUser models.User
	if err := authcontroller.DB.Where("username = ? OR email = ?", request.Username, request.Email).First(&existingUser).Error; err == nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Username or Email already exists")
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	if err := authcontroller.DB.Create(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user")
	}

	return utils.SuccesResponse(c, fiber.StatusCreated, "user succesfull created", user)
}

func (authContoller *AuthController) Login(c *fiber.Ctx) error {
	var request models.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Username and password required")
	}

	var user models.User
	if err := authContoller.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "invalid username credentials")
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "invalid password credential")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	response := fiber.Map{
		"data_user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"token": token,
	}

	return utils.SuccesResponse(c, fiber.StatusOK, "Login Successful", response)
}
