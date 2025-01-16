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

// GetAvatar godoc
// @Summary 获取头像
// @Schemes
// @Description 获取头像
// @Tags 头像模块
// @Produce image/jpeg, image/png, image/webp, image/gif
// @Param hash path string true "email hash"
// @Success 200 {image} image/jpeg,image/png,image/webp,image/gif "avatar"
// @Router /avatar/:hash [get]
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

// UpdateAvatar godoc
// @Summary 修改头像
// @Schemes
// @Description 修改头像
// @Tags 头像模块
// @Accept multipart/form-data
// @Produce json
// @Security Bearer
// @Param file formData file true "avatar"
// @Success 200 {object} v1.Response
// @Router /avatar [put]
func (h *AvatarHandler) UpdateAvatar(ctx *gin.Context) {

	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	req := new(v1.UpdateAvatarRequest)
	if err := HandleUpload(ctx, req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.avatarService.UpdateAvatar(ctx, userId, req); err != nil {
		h.logger.WithContext(ctx).Error("avatarService.UpdateAvatar error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// DeleteAvatar godoc
// @Summary 删除头像
// @Schemes
// @Description 删除头像
// @Tags 头像模块
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.Response
// @Router /avatar [delete]
func (h *AvatarHandler) DeleteAvatar(ctx *gin.Context) {

	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	if err := h.avatarService.DeleteAvatar(ctx, userId); err != nil {
		h.logger.WithContext(ctx).Error("avatarService.DeleteAvatar error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func HandleUpload(ctx *gin.Context, req *v1.UpdateAvatarRequest) error {

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
