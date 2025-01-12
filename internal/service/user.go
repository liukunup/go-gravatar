package service

import (
	"context"
	v1 "go-gravatar/api/v1"
	"go-gravatar/internal/model"
	"go-gravatar/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	ForgotPassword(ctx context.Context, req *v1.ForgotPasswordRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
}

func NewUserService(
	service *Service,
	userRepository repository.UserRepository,
) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
	}
}

type userService struct {
	*Service
	userRepository repository.UserRepository
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {

	// Check email
	user, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if err == nil && user != nil {
		return v1.ErrEmailAlreadyUse
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Generate random userId
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user = &model.User{
		UserId:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Create user
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepository.Create(ctx, user); err != nil {
			return err
		}

		// TODO: send activation email

		return nil
	})
	return err
}

func (s *userService) ForgotPassword(ctx context.Context, req *v1.ForgotPasswordRequest) error {

	// Check email
	user, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if user == nil {
		return v1.ErrEmailNotExists
	}

	// TODO: send reset password email

	return nil
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {

	// Check email
	user, err := s.userRepository.GetByEmail(ctx, req.Username)
	if err != nil {
		return "", v1.ErrInternalServerError
	}
	if user == nil {
		// Check username
		user, err = s.userRepository.GetByUsername(ctx, req.Username)
		if err != nil {
			return "", v1.ErrInternalServerError
		}
		if user == nil {
			return "", v1.ErrUnauthorized
		}
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	// Generate token
	token, err := s.jwt.GenToken(user.Username, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {

	// Get user
	user, err := s.userRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {

	// Get user
	user, err := s.userRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	// Check username
	if req.Username != "" && req.Username != user.Username {
		existingUser, err := s.userRepository.GetByUsername(ctx, req.Username)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return v1.ErrUsernameAlreadyUse
		}

		user.Username = req.Username
	}

	// Check nickname
	if req.Nickname != "" && req.Nickname != user.Nickname {
		user.Nickname = req.Nickname
	}

	// Check email
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := s.userRepository.GetByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return v1.ErrEmailAlreadyUse
		}

		user.Email = req.Email
	}

	if err = s.userRepository.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
