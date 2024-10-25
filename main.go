// main.go
package main

import (
	"log"
	"os"

	"github.com/ABHINAVGARG05/trading-platform/controllers"
	"github.com/ABHINAVGARG05/trading-platform/database"
	"github.com/ABHINAVGARG05/trading-platform/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: No .env file found")
    }

    database.Connect()

    app := fiber.New()

    routes.Setup(app)

    controllers.PriceTicker()
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"  // Set a default port if not provided
    }

    log.Printf("Server running on http://localhost:%s", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }

}
