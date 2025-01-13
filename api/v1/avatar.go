package v1

// 获取头像
type GetAvatarRequest struct {
	Hash string `uri:"hash" binding:"required" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}
type GetAvatarResponseData struct {
	ImageData []byte `json:"imageData"`
	ImageURL  string `json:"imageURL" example:"https://example.com/avatar.png"`
	ImageFile string `json:"imageFile" example:"images/avatar.png"`
	ObjectKey string `json:"objectKey" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}

// 创建头像 or 更新头像
type UpdateAvatarRequest struct {
	UserId string `json:"userId" example:"ExWFdl17WS"`

	ImageData []byte `json:"imageData"`
	ImageURL  string `json:"imageURL" example:"https://example.com/avatar.png"`
	ImageFile string `json:"imageFile" example:"images/avatar.png"`
	ObjectKey string `json:"objectKey" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}

// 删除头像
type DeleteAvatarRequest struct {
	UserId string `json:"userId" example:"ExWFdl17WS"`
}
