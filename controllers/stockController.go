package controllers

import (
	"github.com/ABHINAVGARG05/trading-platform/models"
	"github.com/ABHINAVGARG05/trading-platform/database"
	"github.com/gofiber/fiber/v2"
);

func GetAllStocks(c *fiber.Ctx) error{
	var stocks[] models.Stock
	result := database.DB.Find(&stocks)
	if result.Error != nil{
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(stocks)
}

func GetStockById(c *fiber.Ctx)error{
	id:= c.Params("id")
	var stock models.Stock
	result := database.DB.Find(&stock,id)
	if result.Error != nil{
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(stock)
}