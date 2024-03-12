package handler

import (
	"github.com/gofiber/fiber/v2"
	"golangtest.com/client/auth/dto"
	"golangtest.com/middleware"
	"golangtest.com/models"
	"golangtest.com/utils"
)

func Login(c *fiber.Ctx) error {
	user := models.User{}
	login := dto.Login{}

	if err := c.BodyParser(&login); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Please check your input")
	}

	checkUser := models.DB.Where("email = ?", login.Email).First(&user)

	if checkUser.Error != nil {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := middleware.GenerateJWTToken(user)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error create JWT")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})

}
