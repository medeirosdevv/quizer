package sqlite

import (
    "quiz-backend/internal/domain/entity"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
    var user entity.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
    var user entity.User
    err := r.db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) FindByEmailOrUsername(identifier string) (*entity.User, error) {
    var user entity.User
    err := r.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*entity.User, error) {
    var user entity.User
    err := r.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}