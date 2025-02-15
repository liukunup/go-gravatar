package service_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	v1 "go-gravatar/api/v1"
	"go-gravatar/pkg/jwt"
	mock_repository "go-gravatar/test/mocks/repository"
	"os"
	"testing"

	"go-gravatar/internal/model"
	"go-gravatar/internal/service"
	"go-gravatar/pkg/config"
	"go-gravatar/pkg/log"
	"go-gravatar/pkg/sid"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var (
	logger *log.Logger
	j      *jwt.JWT
	sf     *sid.Sid
)

func TestMain(m *testing.M) {
	fmt.Println("begin")

	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		panic(err)
	}

	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	j = jwt.NewJwt(conf)
	sf = sid.NewSid()

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &v1.RegisterRequest{
		Password: "password",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByEmail(ctx, req.Email).Return(nil, nil)
	mockTm.EXPECT().Transaction(ctx, gomock.Any()).Return(nil)

	err := userService.Register(ctx, req)

	assert.NoError(t, err)
}

func TestUserService_Register_UsernameExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)
	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &v1.RegisterRequest{
		Password: "password",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByEmail(ctx, req.Email).Return(&model.User{}, nil)

	err := userService.Register(ctx, req)

	assert.Error(t, err)
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)
	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &v1.LoginRequest{
		Username: "xxx@gmail.com",
		Password: "password",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Error("failed to hash password")
	}

	mockUserRepo.EXPECT().GetByEmail(ctx, req.Username).Return(&model.User{
		Password: string(hashedPassword),
	}, nil)

	token, err := userService.Login(ctx, req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestUserService_Login_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)
	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &v1.LoginRequest{
		Username: "xxx@gmail.com",
		Password: "password",
	}

	mockUserRepo.EXPECT().GetByEmail(ctx, req.Username).Return(nil, errors.New("user not found"))

	_, err := userService.Login(ctx, req)

	assert.Error(t, err)
}

func TestUserService_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)
	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := "123"

	mockUserRepo.EXPECT().GetByID(ctx, userId).Return(&model.User{
		UserId: userId,
		Email:  "test@example.com",
	}, nil)

	user, err := userService.GetProfile(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, userId, user.UserId)
}

func TestUserService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)
	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := "123"
	req := &v1.UpdateProfileRequest{
		Username: "testuser_1",
		Nickname: "testuser_2",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByID(ctx, userId).Return(&model.User{
		UserId:   userId,
		Username: "old_user_1",
		Nickname: "old_user_2",
		Email:    "old@example.com",
	}, nil)
	mockUserRepo.EXPECT().GetByUsername(ctx, req.Username).Return(nil, v1.ErrNotFound)
	mockUserRepo.EXPECT().GetByEmail(ctx, req.Email).Return(nil, v1.ErrNotFound)
	mockUserRepo.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	mockTm.EXPECT().Transaction(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		}).Return(nil)

	err := userService.UpdateProfile(ctx, userId, req)

	assert.NoError(t, err)
}

func TestUserService_UpdateProfile_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockTm := mock_repository.NewMockTransaction(ctrl)
	srv := service.NewService(mockTm, logger, sf, j)
	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := "123"
	req := &v1.UpdateProfileRequest{
		Username: "test",
		Nickname: "testuser",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByID(ctx, userId).Return(nil, errors.New("user not found"))

	err := userService.UpdateProfile(ctx, userId, req)

	assert.Error(t, err)
}
