// repository/user_repository.go
package repositories

import (
	"ecommerce-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uint) (*models.User, error)
	UpdateUserProfile(user *models.User) error
	DeleteUserProfile(user *models.User) error

}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	// Password should already be hashed by service layer
	return r.db.Create(user).Error
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}
//update current user profile
func (r *userRepository) UpdateUserProfile(user *models.User) error {
	// Update user profile in the database
	return r.db.Save(user).Error
}
// delete current user profile
func (r *userRepository) DeleteUserProfile(user *models.User) error{

	return r.db.Delete(user).Error
}