// main.go
package main

import (
	"ecommerce-api/database"
	"ecommerce-api/handlers"
	
	"ecommerce-api/repositories"
	"ecommerce-api/service"
	"log"
	

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"ecommerce-api/routes"
	"ecommerce-api/config"
	
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	cfg := config.New()

	// Initialize database connection
	database.Connect()

	// Setup dependency injection
	userRepo := repositories.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepo, cfg.JWT_SECRET)

	// Initialize email service
	emailService, err := service.NewEmailService(cfg.EMAIL_API_KEY, cfg.EMAIL_SENDER," ")
	if err != nil {
		log.Fatalf("Failed to initialize email service: %v", err)
	}
	authHandler := handler.NewAuthHandler(authService, emailService)

	orderRepo := repositories.NewOrderRepository(database.DB)
	productRepo := repositories.NewProductRepository(database.DB)
	orderservice := service.NewOrderService(orderRepo,productRepo)
	orderhandler := handler.NewOrderHandler(orderservice)

	productservice := service.NewProductService(productRepo)
	producthandler := handler.NewProductHandler(productservice)

	// Initialize Fiber app
	app := fiber.New()

	routes.SetupRoutes(app, authHandler, producthandler, orderhandler)

	

	// Middleware
	app.Use(cors.New(cors.Config{
	AllowOrigins: "https://booksbyjohn.herokuapp.com, http://localhost:3000",
	AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	// AllowCredentials: true, // Enable if using cookies/auth headers
	// MaxAge: 86400, // Cache CORS preflight for 1 day
}))
	app.Use(logger.New())
	
	// Start server
	log.Fatal(app.Listen(":" + cfg.AppPort))
}