package v1

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"username@example.lan"`
	Password string `json:"password" binding:"required" example:"password"`
}

type ResetRequest struct {
	Email string `json:"email" binding:"required,email" example:"username@example.lan"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"username"`
	Password string `json:"password" binding:"required" example:"password"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Username string `json:"username" example:"username"`
	Nickname string `json:"nickname" example:"Billy"`
	Email    string `json:"email" example:"username@example.lan"`
}
type GetProfileResponseData struct {
	UserId   string `json:"userId" example:"ExWFdl17WS"`
	Username string `json:"username" example:"username"`
	Nickname string `json:"nickname" example:"Billy"`
	Email    string `json:"email" example:"username@example.lan"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
