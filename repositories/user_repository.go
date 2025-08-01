// repository/user_repository.go
package repositories

import (
	"ecommerce-api/models"
	"gorm.io/gorm"
	"time"
)



type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	// Password should already be hashed by service layer
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}
//update current user profile
func (r *UserRepository) UpdateUserProfile(user *models.User) error {
	// Update user profile in the database
	return r.DB.Save(user).Error
}
// delete current user profile
func (r *UserRepository) DeleteUserProfile(user *models.User) error{

	return r.DB.Delete(user).Error
}

// UpdatePassword updates the user's password in the database.
// It takes the user ID and the new password as parameters.
func (r *UserRepository) UpdatePassword(userID int, newPassword string) error {
	
	return r.DB.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}
	
// CreateResetToken creates a new password reset token in the database.
// It takes a PasswordResetToken model as a parameter and returns an error if any occurs.
func (r *UserRepository) CreateResetToken(token *models.PasswordResetToken) error {
	return r.DB.Create(token).Error
}

// FindValidToken checks if a password reset token is valid.
func (r *UserRepository) FindValidToken(token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken
	err := r.DB.Where("token = ? AND used = ? AND expires_at > ?", token, false, time.Now()).First(&resetToken).Error
	return &resetToken, err
}

// MarkTokenAsUsed marks a password reset token as used in the database.
func (r *UserRepository) MarkTokenAsUsed(tokenID uint) error {
	return r.DB.Model(&models.PasswordResetToken{}).Where("id = ?", tokenID).Update("used", true).Error
}