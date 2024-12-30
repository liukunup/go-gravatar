package handler

import (
	"github.com/gin-gonic/gin"
	"go-gravatar/internal/service"
)

type AvatarHandler struct {
	*Handler
	avatarService service.AvatarService
}

func NewAvatarHandler(
    handler *Handler,
    avatarService service.AvatarService,
) *AvatarHandler {
	return &AvatarHandler{
		Handler:      handler,
		avatarService: avatarService,
	}
}

func (h *AvatarHandler) GetAvatar(ctx *gin.Context) {

}
