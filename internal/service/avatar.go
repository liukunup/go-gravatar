package service

import (
    "context"
	"go-gravatar/internal/model"
	"go-gravatar/internal/repository"
)

type AvatarService interface {
	GetAvatar(ctx context.Context, id int64) (*model.Avatar, error)
}
func NewAvatarService(
    service *Service,
    avatarRepository repository.AvatarRepository,
) AvatarService {
	return &avatarService{
		Service:        service,
		avatarRepository: avatarRepository,
	}
}

type avatarService struct {
	*Service
	avatarRepository repository.AvatarRepository
}

func (s *avatarService) GetAvatar(ctx context.Context, id int64) (*model.Avatar, error) {
	return s.avatarRepository.GetAvatar(ctx, id)
}
