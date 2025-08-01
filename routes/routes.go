// routes/routes.go
package routes

import (
	"ecommerce-api/handlers"
	"ecommerce-api/middleware"
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App, 
	authHandler *handler.AuthHandler,
	productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler) {

	// Public routes (no auth required)
	public := app.Group("/api")
	public.Post("/register", authHandler.Register)
	public.Post("/login", authHandler.Login)
	// Forgot password route
	public.Post("/reset-password", authHandler.ResetPasswordHandler) // Reset password route
	public.Get("/products", productHandler.GetProducts)
	public.Get("/products/:id", productHandler.GetProduct)

	// Protected routes (require valid JWT)
	protected := app.Group("/api")
	protected.Use(middleware.JWTProtected(authHandler.AuthService.UserRepo)) // Apply JWT middleware to all routes below
	
	// Product routes (admin only)
	productAdmin := protected.Group("/products")
	productAdmin.Use(middleware.AdminOnly()) // Additional admin check
	productAdmin.Post("/", productHandler.CreateProduct)
	productAdmin.Put("/:id", productHandler.UpdateProduct)
	productAdmin.Delete("/:id", productHandler.DeleteProduct)

	// Order routes (authenticated users)
	order := protected.Group("/orders")
	order.Post("/", orderHandler.CreateOrder)
	order.Get("/", orderHandler.GetUserOrders)
	order.Get("/:id", orderHandler.GetOrderByID)
	order.Delete("/:id", orderHandler.CancelOrder)

	// User profile routes
	user := protected.Group("/users")
	user.Use(middleware.JWTProtected(authHandler.AuthService.UserRepo)) // Ensure user is authenticated
	user.Get("/me", authHandler.GetCurrentUser) // Get current user profile
	user.Put("/me", authHandler.UpdateUserProfile) // Update user profile
	// delete current account
	user.Delete("/me", authHandler.DeleteUserProfile) // Delete user profile
	// Add user profile routes here when implemented
}

// Example middleware package content:
// middleware/auth.go
