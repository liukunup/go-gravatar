package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrUsernameAlreadyUse = newError(1001, "The username is already in use.")
	ErrEmailAlreadyUse    = newError(1002, "The email is already in use.")
	ErrEmailNotExists     = newError(1003, "The email does not exist.")
)
