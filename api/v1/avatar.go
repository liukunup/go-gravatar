package v1

type GetAvatarRequest struct {
	Hash string `uri:"hash" binding:"required" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}
type GetAvatarResponseData struct {
	ImageData []byte `json:"imageData"`
	ImageURL  string `json:"imageURL" example:"https://example.com/avatar.png"`
	ImageFile string `json:"imageFile" example:"images/avatar.png"`
	ObjectKey string `json:"objectKey" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}
type GetAvatarResponse struct {
	Response
	Data GetAvatarResponseData
}

type UpdateAvatarRequest struct {
	ImageData []byte `json:"imageData"`
	ImageURL  string `json:"imageURL" example:"https://example.com/avatar.png"`
	ImageFile string `json:"imageFile" example:"images/avatar.png"`
	ObjectKey string `json:"objectKey" example:"e1e7ba949ade0936e071132d2edd3b3c"`
}
