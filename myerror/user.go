package myerror

import (
	"fmt"
	"myapp/apperror"
	"net/http"
)

func ErrUserGet(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10000",
		Message:   "Failed to get example.",
		IsSentry:  true,
	}
}

func ErrUserCreate(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10001",
		Message:   "Failed to create example.",
		IsSentry:  true,
	}
}

func ErrUserUpdate(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10002",
		Message:   "Failed to update example.",
		IsSentry:  true,
	}
}

func ErrUserDelete(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10003",
		Message:   "Failed to delete example.",
		IsSentry:  true,
	}
}

func ErrUserNotFound() apperror.TeqError {
	return apperror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrUserInvalidParam(param string) apperror.TeqError {
	return apperror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrCommitTransaction(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "002",
		Message:   "Failed to commit transaction",
		IsSentry:  true,
	}
}

func ErrUnauthorized(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "000001",
		Message:   "Unauthorized!",
		IsSentry:  false,
	}
}

func ErrNoPermission() apperror.TeqError {
	return apperror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusForbidden,
		ErrorCode: "000002",
		Message:   "No permission.",
		IsSentry:  false,
	}
}

func ErrInvalidParams(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "000003",
		Message:   "Invalid params.",
		IsSentry:  false,
	}
}

func ErrNotFound(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "000004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrGet(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000005",
		Message:   "Failed to get.",
		IsSentry:  true,
	}
}

func ErrCreate(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000006",
		Message:   "Failed to create.",
		IsSentry:  true,
	}
}

func ErrUpdate(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000007",
		Message:   "Failed to update.",
		IsSentry:  true,
	}
}

func ErrDelete(err error) apperror.TeqError {
	return apperror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000008",
		Message:   "Failed to delete.",
		IsSentry:  true,
	}
}
