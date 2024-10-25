package routes
import (
	"github.com/gofiber/fiber/v2"
	"github.com/ABHINAVGARG05/trading-platform/controllers"
)

func Setup(app *fiber.App){
	api := app.Group("api/v1")
	api.Post("/user/register",controllers.RegisterUser)
	api.Get("/stocks",controllers.GetAllStocks)
	api.Get("/stock/:id",controllers.GetStockById)
	api.Post("/stock/buy",controllers.BuyStock)
	api.Post("/stock/sell",controllers.SellStock)
	api.Get("/user/:id",controllers.GetUserInfo)
}