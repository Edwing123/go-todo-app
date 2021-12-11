package models

import "errors"

var (
	ErrDuplicatedUsername = errors.New("models: duplicated username")
)
