package v1

// 注册
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"username@example.lan"`
	Password string `json:"password" binding:"required" example:"password"`
}

// 忘记密码
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email" example:"username@example.lan"`
}

// 登录
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

// 更新信息
type UpdateProfileRequest struct {
	Id       uint   `json:"id" binding:"required,id" example:"1"`
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
