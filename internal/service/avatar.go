package service

import (
	"context"
	v1 "go-gravatar/api/v1"
	"go-gravatar/internal/model"
	"go-gravatar/internal/repository"
)

type AvatarService interface {
	GetAvatar(ctx context.Context, req *v1.GetAvatarRequest) (*v1.GetAvatarResponseData, error)
	CreateOrUpdateAvatar(ctx context.Context, req *v1.CreateOrUpdateAvatarRequest) error
	DeleteAvatar(ctx context.Context, req *v1.DeleteAvatarRequest) error
}

func NewAvatarService(
	service *Service,
	avatarRepository repository.AvatarRepository,
) AvatarService {
	return &avatarService{
		Service:          service,
		avatarRepository: avatarRepository,
	}
}

type avatarService struct {
	*Service
	avatarRepository repository.AvatarRepository
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

func (s *avatarService) CreateOrUpdateAvatar(ctx context.Context, req *v1.CreateOrUpdateAvatarRequest) error {

	avatar, err := s.avatarRepository.GetByHash(ctx, req.Hash)
	if err != nil {
		avatar = &model.Avatar{
			Hash:      req.Hash,
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
		avatar.ImageData = req.ImageData
		avatar.ImageURL = req.ImageURL
		avatar.ImageFile = req.ImageFile
		avatar.ObjectKey = req.ObjectKey

		err = s.avatarRepository.Update(ctx, avatar)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *avatarService) DeleteAvatar(ctx context.Context, req *v1.DeleteAvatarRequest) error {

	err := s.avatarRepository.Delete(ctx, req.Hash)
	if err != nil {
		return err
	}

	return nil
}
