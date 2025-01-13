package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	v1 "go-gravatar/api/v1"
	"go-gravatar/internal/model"
	"go-gravatar/internal/repository"
	"strings"
)

type AvatarService interface {
	GetAvatar(ctx context.Context, req *v1.GetAvatarRequest) (*v1.GetAvatarResponseData, error)
	UpdateAvatar(ctx context.Context, req *v1.UpdateAvatarRequest) error
	DeleteAvatar(ctx context.Context, req *v1.DeleteAvatarRequest) error
}

func NewAvatarService(
	service *Service,
	avatarRepository repository.AvatarRepository,
	userRepository repository.UserRepository,
) AvatarService {
	return &avatarService{
		Service:          service,
		avatarRepository: avatarRepository,
		userRepository:   userRepository,
	}
}

type avatarService struct {
	*Service
	avatarRepository repository.AvatarRepository
	userRepository   repository.UserRepository
}

func (s *avatarService) GetAvatar(ctx context.Context, req *v1.GetAvatarRequest) (*v1.GetAvatarResponseData, error) {

	avatar, err := s.avatarRepository.GetByHash(ctx, req.Hash)
	if err != nil {
		return nil, err
	}

	return &v1.GetAvatarResponseData{
		ImageData: avatar.ImageData,
		ImageURL:  avatar.ImageURL,
		ImageFile: avatar.ImageFile,
		ObjectKey: avatar.ObjectKey,
	}, nil
}

func (s *avatarService) UpdateAvatar(ctx context.Context, req *v1.UpdateAvatarRequest) error {

	user, err := s.userRepository.GetByID(ctx, req.UserId)
	if err != nil {
		return err
	}

	email := strings.TrimSpace(strings.ToLower(user.Email))
	hash := md5.New()
	hash.Write([]byte(email))
	hashedEmail := hex.EncodeToString(hash.Sum(nil))

	avatar, err := s.avatarRepository.GetByHash(ctx, hashedEmail)
	if err == v1.ErrNotFound && avatar == nil {
		avatar = &model.Avatar{
			Hash:      hashedEmail,
			ImageData: req.ImageData,
			ImageURL:  req.ImageURL,
			ImageFile: req.ImageFile,
			ObjectKey: req.ObjectKey,
		}
		err := s.avatarRepository.Create(ctx, avatar)
		if err != nil {
			return err
		}
	} else {
		if err == nil && avatar != nil {

			if req.ImageData != nil {
				avatar.ImageData = req.ImageData
			}
			if req.ImageURL != "" {
				avatar.ImageURL = req.ImageURL
			}
			if req.ImageFile != "" {
				avatar.ImageFile = req.ImageFile
			}
			if req.ObjectKey != "" {
				avatar.ObjectKey = req.ObjectKey
			}

			err := s.avatarRepository.Update(ctx, avatar)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *avatarService) DeleteAvatar(ctx context.Context, req *v1.DeleteAvatarRequest) error {

	user, err := s.userRepository.GetByID(ctx, req.UserId)
	if err != nil {
		return err
	}

	email := strings.TrimSpace(strings.ToLower(user.Email))
	hash := md5.New()
	hash.Write([]byte(email))
	hashedEmail := hex.EncodeToString(hash.Sum(nil))

	err = s.avatarRepository.Delete(ctx, hashedEmail)
	if err != nil {
		return err
	}

	return nil
}
