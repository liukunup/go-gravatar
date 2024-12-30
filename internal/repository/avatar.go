package repository

import (
    "context"
	"go-gravatar/internal/model"
)

type AvatarRepository interface {
	GetAvatar(ctx context.Context, id int64) (*model.Avatar, error)
}

func NewAvatarRepository(
	repository *Repository,
) AvatarRepository {
	return &avatarRepository{
		Repository: repository,
	}
}

type avatarRepository struct {
	*Repository
}

func (r *avatarRepository) GetAvatar(ctx context.Context, id int64) (*model.Avatar, error) {
	var avatar model.Avatar

	return &avatar, nil
}
