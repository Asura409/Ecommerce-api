// handler/order_handler.go
package handler

import (
	"ecommerce-api/models"
	"ecommerce-api/service"
	"github.com/gofiber/fiber/v2"
)


type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	
	order.UserID = user.ID
	order.Status = "pending"
	
	if err := h.OrderService.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(order)
}

// Implement other handlers...
func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	productID := uint(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}
	user := c.Locals("user").(*models.User)

	order, err := h.OrderService.GetOrder(productID, user.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	return c.JSON(order)
}

func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	
	orders, err := h.OrderService.GetUserOrders(user.ID, 1, 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve orders",
		})
	}
	
	return c.JSON(orders)
}

func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	productID := uint(id)
	if err !=	 nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	user := c.Locals("user").(*models.User)	
	order, err := h.OrderService.GetOrder(productID , user.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}
	if order.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are not authorized to cancel this order",
		})
	}

	if err := h.OrderService.CancelOrder(productID, user.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not cancel order",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

