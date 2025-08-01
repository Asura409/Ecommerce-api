// service/order_service.go
package service
import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"errors"
)




type OrderService struct {
	OrderRepo    *repositories.OrderRepository
	ProductRepo  *repositories.ProductRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, productRepo *repositories.ProductRepository) *OrderService {
	return &OrderService{
		OrderRepo: orderRepo,
		ProductRepo: productRepo,
	}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	// Validate stock availability
	for _, item := range order.OrderItems {
		product, err := s.ProductRepo.FindByID(item.ProductID)
		if err != nil {
			return err
		}
		if product.StockQuantity < item.Quantity {
			return errors.New("insufficient stock")
		}
	}
	
	// Calculate total amount
	var total float64
	for _, item := range order.OrderItems {
		product, _ := s.ProductRepo.FindByID(item.ProductID)
		total += product.Price * float64(item.Quantity)
		
	}
	order.TotalAmount = total
	
	return s.OrderRepo.Create(order)
}

// Implement other methods...
func (s *OrderService) GetOrder(id uint, userID uint) (*models.Order, error) {
	order, err := s.OrderRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if order.UserID != userID {
		return nil, errors.New("unauthorized access")
	}
	return order, nil
}

func (s *OrderService) GetUserOrders(userID uint, page, limit int) ([]models.Order, error) {
	orders, err := s.OrderRepo.FindByUserID(userID, page, limit)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) CancelOrder(id uint, userID uint) error {
	order, err := s.OrderRepo.FindByID(id)
	if err != nil {
		return err
	}
	if order.UserID != userID {
		return errors.New("unauthorized access")
	}
	if order.Status != "pending" {
		return errors.New("only pending orders can be canceled")
	}
	
	order.Status = "canceled"
	return s.OrderRepo.UpdateStatus(id, order.Status)
}