package models

import (
	"gorm.io/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
)

// User represents the user model
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:100;unique;not null" json:"email"`
	Password  string         `gorm:"size:100;not null" json:"-"`
	Role      string         `gorm:"size:20;default:'customer'" json:"role"`
}

// Product represents the product model
type Product struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"size:255;not null" json:"name"`
	Description  string         `json:"description"`
	Price        float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Category     string         `gorm:"size:100" json:"category"`
	StockQuantity int           `gorm:"not null" json:"stock_quantity"`
	ImageURL     string         `gorm:"size:255" json:"image_url"`
}

// Order represents the order model
type Order struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	UserID         uint           `json:"user_id"`
	User           User           `gorm:"foreignKey:UserID" json:"-"`
	TotalAmount    float64        `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status         string         `gorm:"size:50;default:'pending'" json:"status"`
	ShippingAddress string        `gorm:"type:text;not null" json:"shipping_address"`
	PaymentMethod  string         `gorm:"size:50;not null" json:"payment_method"`
	OrderItems     []OrderItem    `gorm:"foreignKey:OrderID" json:"items"`
}

// OrderItem represents items within an order
type OrderItem struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	OrderID        uint           `json:"order_id"`
	ProductID      uint           `json:"product_id"`
	Product        Product        `gorm:"foreignKey:ProductID" json:"product"`
	Quantity       int            `gorm:"not null" json:"quantity"`
	PriceAtPurchase float64       `gorm:"type:decimal(10,2);not null" json:"price_at_purchase"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided password with the stored hash
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}