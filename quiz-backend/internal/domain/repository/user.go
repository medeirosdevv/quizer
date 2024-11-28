package repository

import "quiz-backend/internal/domain/entity"

type UserRepository interface {
    Create(user *entity.User) error
    FindByEmail(email string) (*entity.User, error)
    FindByUsername(username string) (*entity.User, error)
    FindByEmailOrUsername(identifier string) (*entity.User, error)
    FindByID(id uint) (*entity.User, error) 
}