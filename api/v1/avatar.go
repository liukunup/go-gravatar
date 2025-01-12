package v1

// 获取头像
type GetAvatarRequest struct {
	Hash string `uri:"hash" binding:"required,hash" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}
type GetAvatarResponseData struct {
	ImageData []byte `json:"imageData" example:"[0,1,2,3,4,5,6,7,8,9]"`
	ImageURL  string `json:"imageURL" example:"https://example.com/avatar.png"`
	ImageFile string `json:"imageFile" example:"images/avatar.png"`
	ObjectKey string `json:"objectKey" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}

// 创建头像 更新头像
type CreateOrUpdateAvatarRequest struct {
	Hash string `uri:"hash" binding:"required,hash" example:"e1e7ba949ade0936e071132d2edd3b3c"`

	ImageData []byte `json:"imageData" example:"[0,1,2,3,4,5,6,7,8,9]"`
	ImageURL  string `json:"imageURL" example:"https://example.com/avatar.png"`
	ImageFile string `json:"imageFile" example:"images/avatar.png"`
	ObjectKey string `json:"objectKey" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}

// 删除头像
type DeleteAvatarRequest struct {
	Hash string `uri:"hash" binding:"required,hash" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}
