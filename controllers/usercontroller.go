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
