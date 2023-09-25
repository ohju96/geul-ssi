package custom

import "net/http"

var (

	// 1000 ~ 1999 : common
	BadRequest          = NewError(http.StatusBadRequest, 1001, "bad request")
	NotFound            = NewError(http.StatusNotFound, 1002, "not found")
	TxGenerateFail      = NewError(http.StatusInternalServerError, 1004, "tx generate fail")
	TxRollbackFail      = NewError(http.StatusInternalServerError, 1005, "tx rollback fail")
	TxCommitFail        = NewError(http.StatusInternalServerError, 1006, "tx commit fail")
	RequestBindFail     = NewError(http.StatusBadRequest, 1007, "request bind fail")
	RequestValidateFail = NewError(http.StatusBadRequest, 1008, "request validate fail")

	// 2000 ~ 2999 : utils
	InternalServerError = NewError(http.StatusInternalServerError, 1003, "internal server error")
	BcryptGenerateFail  = NewError(http.StatusInternalServerError, 2001, "bcrypt generate fail")
	BcryptParsingFail   = NewError(http.StatusInternalServerError, 2002, "bcrypt parsing fail")
	BcryptMismatch      = NewError(http.StatusUnauthorized, 2003, "bcrypt mismatch")
	UUIDGenerateFail    = NewError(http.StatusInternalServerError, 2004, "uuid generate fail")

	// 3000 ~ 3999 : user
	SamePassword   = NewError(http.StatusUnauthorized, 3003, "same password")
	AlreadyExist   = NewError(http.StatusConflict, 3004, "already exist")
	Unauthorized   = NewError(http.StatusUnauthorized, 3005, "unauthorized")
	PasswdNotMatch = NewError(http.StatusUnauthorized, 3006, "password not match")
)

type Error struct {
	StatusCode int
	Code       int
	Msg        string
}

func NewError(statusCode int, code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
	}
}
