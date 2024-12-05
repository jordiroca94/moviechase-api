package user

import (
	"fmt"
	"time"

	"github.com/jordiroca94/moviechase-api/service/auth"
	"github.com/jordiroca94/moviechase-api/types"
)

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUserByEmail(email string) (*types.User, error) {
	u, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *UserService) CreateToken(secret []byte, id int, email string, firstName string, lastName string) (string, error) {
	token, err := auth.CreateJWT(secret, id, email, firstName, lastName)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	return token, nil
}

func (s *UserService) CreateUser(user types.RegisterUserPayload, hashedPassword string) error {
	err := s.repository.CreateUser(&types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("user not found")
	}
	return nil
}
