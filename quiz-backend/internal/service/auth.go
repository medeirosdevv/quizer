package service

import (
	"errors"
	"quiz-backend/internal/domain/entity"
	"quiz-backend/internal/domain/repository"
	"quiz-backend/pkg/jwt"
)

type AuthService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func (s *AuthService) Register(input RegisterInput) error {
	if _, err := s.userRepo.FindByEmail(input.Email); err == nil {
		return errors.New("email já cadastrado")
	}

	if _, err := s.userRepo.FindByUsername(input.Username); err == nil {
		return errors.New("nome de usuário já existe")
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Username, input.Password)
	if err != nil {
		return err
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(input LoginInput) (string, *entity.User, error) {
	user, err := s.userRepo.FindByEmailOrUsername(input.Identifier)
	if err != nil {
		return "", nil, errors.New("usuário não encontrado")
	}

	if err := user.CheckPassword(input.Password); err != nil {
		return "", nil, errors.New("senha incorreta")
	}

	token, err := jwt.GenerateToken(user.ID, user.Username, s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
