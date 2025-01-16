package service

import (
	"context"
	v1 "go-gravatar/api/v1"
	"go-gravatar/internal/model"
	"go-gravatar/internal/repository"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Reset(ctx context.Context, req *v1.ResetRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	Delete(ctx context.Context, userId string) error
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
	if err != nil && err != v1.ErrNotFound {
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
	// Extract username from email
	re := regexp.MustCompile(`^[^@]+`)
	username := re.FindString(req.Email)
	user = &model.User{
		UserId:   userId,
		Username: username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err_t := s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.userRepository.Create(ctx, user); err != nil {
			return err
		}

		// TODO: send account activation email

		return nil
	})
	return err_t
}

func (s *userService) Reset(ctx context.Context, req *v1.ResetRequest) error {

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
	if err != nil && err != v1.ErrNotFound {
		return "", v1.ErrInternalServerError
	}
	if user == nil {
		// Check username
		user, err = s.userRepository.GetByUsername(ctx, req.Username)
		if err != nil && err != v1.ErrNotFound {
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
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {

	user, err := s.userRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserId,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {

	user, err := s.userRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	// Check username
	if req.Username != "" && req.Username != user.Username {
		existingUser, err := s.userRepository.GetByUsername(ctx, req.Username)
		if err != nil && err != v1.ErrNotFound {
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
		if err != nil && err != v1.ErrNotFound {
			return err
		}
		if existingUser != nil {
			return v1.ErrEmailAlreadyUse
		}

		user.Email = req.Email
	}

	err_t := s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.userRepository.Update(ctx, user); err != nil {
			return err
		}
		return nil
	})
	return err_t
}

func (s *userService) Delete(ctx context.Context, userId string) error {
	err_t := s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.userRepository.Delete(ctx, userId); err != nil {
			return err
		}
		return nil
	})
	return err_t
}
