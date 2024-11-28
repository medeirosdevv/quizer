package service

import (
    "quiz-backend/internal/domain/entity"
    "quiz-backend/internal/domain/repository"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

func (s *UserService) GetUserByID(id uint) (*entity.User, error) {
    return s.userRepo.FindByID(id)
}