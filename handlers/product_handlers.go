// handler/product_handler.go
package handler
import (
	"ecommerce-api/models"
	"ecommerce-api/service"
	"github.com/gofiber/fiber/v2"
)


type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := h.ProductService.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// Implement other handlers similarly...
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := h.ProductService.UpdateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update product",
		})
	}

	return c.JSON(product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	if err := h.ProductService.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	product, err := h.ProductService.GetProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(product)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if  page < 1 {
		page = 1
	}

	limit := c.QueryInt("limit", 10)
	if  limit < 1 {
		limit = 10
	}

	products, total, err := h.ProductService.GetProducts(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve products",
		})
	}

	return c.JSON(fiber.Map{
		"products": products,
		"total":    total,
	})
}
