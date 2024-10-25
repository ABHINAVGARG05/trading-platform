package controllers

import (
	"log"

	"github.com/ABHINAVGARG05/trading-platform/database"
	"github.com/ABHINAVGARG05/trading-platform/models"
	"github.com/gofiber/fiber/v2"
)

func BuyStock(c *fiber.Ctx) error {
	type BuyReq struct {
		UserID   uint `json:"user_id"`
		StockID  uint `json:"stock_id"`
		Quantity uint `json:"quantity"`
	}
	var req BuyReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var user models.User
	var stock models.Stock
	database.DB.First(&user, req.UserID)
	database.DB.First(&stock, req.StockID)
	log.Println(req.UserID)
	log.Println(req.StockID)
	if stock.ID == 0 {
		return c.Status(400).SendString("Stock Not Found")
	}
	cost := float64(req.Quantity) * stock.Price
	if user.Balance < cost {
		return c.Status(400).SendString("Insufficient Balance")
	}

	user.Balance -= cost
	database.DB.Save(&user)
	stock.Demand += req.Quantity
	stock.Volume += req.Quantity
	database.DB.Save(&stock)

	order := models.Order{
		UserID:   req.UserID,
		StockID:  req.StockID,
		Type:     "buy",
		Quantity: req.Quantity,
		Price:    stock.Price,
	}
	//PriceUpdater()
	database.DB.Create(&order)

	var portfolio models.Portfolio
	if err := database.DB.Where("user_id = ? AND stock_id = ?", req.UserID, req.StockID).First(&portfolio).Error; err != nil {
        // Portfolio doesn't exist; create a new one
        portfolio = models.Portfolio{
            UserID:   req.UserID,
            StockID:  req.StockID,
            Quantity: req.Quantity,
            AvgPrice: stock.Price,
        }
        database.DB.Create(&portfolio)
        log.Println("Portfolio created")
    } else {
		totalQuantity := portfolio.Quantity + req.Quantity
		portfolio.AvgPrice = (portfolio.AvgPrice*float64(portfolio.Quantity) + stock.Price*float64(req.Quantity)) / float64(totalQuantity)
		portfolio.Quantity += totalQuantity
		database.DB.Save(&portfolio)
		log.Println("portfolio updated")
	}
	return c.JSON(order)
}

func SellStock(c *fiber.Ctx) error {
	type SellReq struct {
		UserId   uint `json:"user_id"`
		StockId  uint `json:"stock_id"`
		Quantity uint `json:"quantity"`
	}
	var req SellReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var user models.User
	var stock models.Stock
	var portfolio models.Portfolio

	database.DB.First(&user, req.UserId)
	database.DB.First(&stock, req.StockId)
	if stock.ID == 0 {
		return c.Status(400).SendString("Stock Not Found")
	}
	database.DB.Where("user_id = ? AND stock_id = ?", req.UserId, req.StockId).First(&portfolio)
	if portfolio.Quantity == 0 || portfolio.Quantity < req.Quantity {
		return c.Status(400).SendString("Insufficient Stock Quantity")
	}

	earnings := float64(req.Quantity) * stock.Price
	user.Balance += earnings
	database.DB.Save(&user)

	stock.Supply += req.Quantity
	stock.Volume -= req.Quantity
	//PriceUpdater()
	database.DB.Save(&stock)

	order := models.Order{
		UserID:   req.UserId,
		StockID:  req.StockId,
		Type:     "sell",
		Quantity: req.Quantity,
		Price:    stock.Price,
	}
	database.DB.Create(&order)
	//database.DB.Where("user_id = ? AND stock_id = ?", req.UserID, req.StockID).First(&portfolio)

	portfolio.Quantity -= req.Quantity
	if portfolio.Quantity == 0 {
		database.DB.Delete(&portfolio)
	} else {
		database.DB.Save(&portfolio)
	}
	return c.JSON(order)
}
