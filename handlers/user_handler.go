// handler/auth_handler.go
package handler

import (
	"ecommerce-api/models"
	"ecommerce-api/service"
	"fmt"
	"time"
	"strings"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *service.AuthService
	EmailSwervice *service.EmailService
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func NewAuthHandler(authService *service.AuthService, emailService *service.EmailService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		EmailSwervice: emailService,
	}
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.AuthService.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	user, token, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

//get current user profile
func (h *AuthHandler) GetCurrentUser(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}
// Update user profile
func (h *AuthHandler) UpdateUserProfile(c *fiber.Ctx) error {	
	user, ok := c.Locals("user").(*models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password // Password should be hashed in service layer

	if err := h.AuthService.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update user",
		})
	}

	return c.JSON(user)
}

//delete user profile
func (h *AuthHandler) DeleteUserProfile(c *fiber.Ctx) error{
	user, ok := c.Locals("user").(*models.User)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	fmt.Printf("Are you sure you want to delete account %v ?",user.Name)

     return h.AuthService.DeleteUserProfile(user)
}

func(h *AuthHandler) ResetPasswordHandler(c *fiber.Ctx) error {
	
		// Parse request
		req := new(ForgotPasswordRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		// Validate email format
		if !isValidEmail(req.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Please enter a valid email address",
			}) // :cite[5]
		}

		// Check if user exists
		var user models.User
		if err := h.AuthService.UserRepo.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Don't reveal whether user exists for security
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "If the email exists, a reset link has been sent",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Database error",
			})
		}

		// Generate reset token (JWT)
		token, err := h.AuthService.GenerateToken(&user)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate token",
			})
		}

		

		// Store token in database (or cache)
		resetToken := models.PasswordResetToken{
			Email:     user.Email,
			Token:     token,
			ExpiresAt: time.Now().Add(time.Hour * 1),
		}
		if err := h.AuthService.UserRepo.DB.Create(&resetToken).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save token",
			})
		}

		// Send email (in production you'd use a mail service)
		go sendResetEmail(user.Email, token)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "If the email exists, a reset link has been sent",
		})
	}


func isValidEmail(email string) bool {
	// Simple email validation - implement proper validation
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func sendResetEmail(to, token string) {
	// Implement actual email sending logic
	// This would typically use SMTP or a service like SendGrid
	// For testing, you might just log it
	fmt.Printf("Reset link for %s: http://yourapp.com/reset-password?token=%s\n", to, token)
}