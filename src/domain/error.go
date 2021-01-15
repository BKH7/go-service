package domain

import (
	"errors"
)

var (
	// ErrInternalServerError = 500
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound = 404
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict = 409
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput = 400
	ErrBadParamInput = errors.New("Given Param is not valid")
	// ErrUnauthorized = 401
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrForbidden = 403
	ErrForbidden = errors.New("Forbidden")
)
