package models

import "errors"

var (
	ErrDuplicatedUsername = errors.New("models: duplicated username")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
)
