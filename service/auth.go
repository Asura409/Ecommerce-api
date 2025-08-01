// service/auth_service.go
package service

import (
	"ecommerce-api/models"
	"ecommerce-api/repositories"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)



type AuthService struct {
	UserRepo repositories.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(user *models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	
	// Set default role if not specified
	if user.Role == "" {
		user.Role = "customer"
	}

	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	user, err := s.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) UpdateUser(user *models.User) error{
	return s.UserRepo.UpdateUserProfile(user)
}
// delete user profile
func (s *AuthService) DeleteUserProfile(user *models.User)error{
	return s.UserRepo.DeleteUserProfile(user)
}

