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
		return "", fmt.Errorf("error creating token")
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
		//A DEFAULT IMAGE
		Image: "https://res.cloudinary.com/dgauzk80l/image/upload/v1743001082/profileplaceholder.png",
	})
	if err != nil {
		return fmt.Errorf("error creating user")
	}
	return nil
}

func (s *UserService) GetUserByID(id int) (*types.User, error) {

	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, user types.UpdateUserPayload) error {
	updatePayload := &types.UpdateUserPayload{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	err := s.repository.UpdateUser(id, updatePayload)
	if err != nil {
		return fmt.Errorf("error updating user")
	}

	return nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("error deleting user")
	}
	return nil
}

func (s *UserService) UpdateUserImage(id int, image string) error {
	err := s.repository.UpdateUserImage(id, image)
	if err != nil {
		return fmt.Errorf("error updating user image")
	}
	return nil
}
