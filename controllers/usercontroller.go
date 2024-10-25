package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/ABHINAVGARG05/trading-platform/database"
    "github.com/ABHINAVGARG05/trading-platform/models"
)

func RegisterUser(c *fiber.Ctx) error {
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    database.DB.Create(&user)
    return c.JSON(user)
}

func GetUserInfo(c* fiber.Ctx)error{
	userId := c.Params("id")
	var user models.User
    if err := database.DB.First(&user, userId).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }
	var portfolio []models.Portfolio
    if err := database.DB.Where("user_id = ?", user.ID).Find(&portfolio).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve portfolio"})
    }

    // Create the response object
    userPortfolio := models.UserPortfolio{
        User:      user,
        Portfolio: portfolio,
    }
	return c.Status(fiber.StatusOK).JSON(userPortfolio)
}
