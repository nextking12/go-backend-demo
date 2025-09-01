package services

import (
	"go-backend-demo/internal/models"
	"go-backend-demo/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return s.userRepo.Create(req)
}

func (s *UserService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return s.userRepo.Update(id, req)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}
