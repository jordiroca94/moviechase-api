package user

import (
	"fmt"

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
