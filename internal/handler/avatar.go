package handler

import (
	v1 "go-gravatar/api/v1"
	"go-gravatar/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		Handler:       handler,
		avatarService: avatarService,
	}
}

func (h *AvatarHandler) GetAvatar(ctx *gin.Context) {

	req := new(v1.GetAvatarRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	avatar, err := h.avatarService.GetAvatar(ctx, req)
	if err != nil {
		h.logger.WithContext(ctx).Error("avatarService.GetAvatar error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	HandleDownload(ctx, avatar)
}

func (h *AvatarHandler) CreateOrUpdateAvatar(ctx *gin.Context) {

	req := new(v1.CreateOrUpdateAvatarRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := HandleUpload(ctx, req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.avatarService.CreateOrUpdateAvatar(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("avatarService.CreateOrUpdateAvatar error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func HandleUpload(ctx *gin.Context, req *v1.CreateOrUpdateAvatarRequest) error {

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	var tempPath = "images/"
	req.ImageFile = tempPath + fileHeader.Filename
	err = ctx.SaveUploadedFile(fileHeader, req.ImageFile)
	if err != nil {
		return err
	}

	return nil
}

func HandleDownload(ctx *gin.Context, avatar *v1.GetAvatarResponseData) {

	if avatar == nil {
		v1.HandleError(ctx, http.StatusNotFound, v1.ErrNotFound, nil)
		return
	}

	if len(avatar.ImageData) > 0 {
		ctx.Data(http.StatusOK, "image/png", []byte(avatar.ImageData))
		return
	}

	if avatar.ImageFile != "" {
		ctx.File(avatar.ImageFile)
		return
	}

	if avatar.ImageURL != "" {
		ctx.Redirect(http.StatusFound, avatar.ImageURL)
		return
	}

	if avatar.ObjectKey != "" {
		// 从对象存储下载
	}
}
