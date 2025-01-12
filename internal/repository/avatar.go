package repository

import (
	"context"
	"errors"
	v1 "go-gravatar/api/v1"
	"go-gravatar/internal/model"

	"gorm.io/gorm"
)

type AvatarRepository interface {
	GetByHash(ctx context.Context, hash string) (*model.Avatar, error)
	Create(ctx context.Context, avatar *model.Avatar) error
	Update(ctx context.Context, avatar *model.Avatar) error
	Delete(ctx context.Context, hash string) error
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

func (r *avatarRepository) GetByHash(ctx context.Context, hash string) (*model.Avatar, error) {
	var avatar model.Avatar
	if err := r.DB(ctx).Where("hash = ?", hash).Where("deleted_at IS NULL").First(&avatar).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &avatar, nil
}

func (r *avatarRepository) Create(ctx context.Context, avatar *model.Avatar) error {
	if err := r.DB(ctx).Create(avatar).Error; err != nil {
		return err
	}
	return nil
}

func (r *avatarRepository) Update(ctx context.Context, avatar *model.Avatar) error {
	if err := r.DB(ctx).Save(avatar).Error; err != nil {
		return err
	}
	return nil
}

func (r *avatarRepository) Delete(ctx context.Context, hash string) error {
	if err := r.DB(ctx).Where("hash = ?", hash).Delete(&model.Avatar{}).Error; err != nil {
		return err
	}
	return nil
}
