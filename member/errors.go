package member

import "errors"

var (
	CONFLICT_ERROR = errors.New("Username already exists")
)
