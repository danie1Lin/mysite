package models

import "mysite/pkg/errors"

const (
	RecordNotFoundErrorCode string = "record_not_found_error"
)

var (
	RecordNotFoundError error = errors.New(RecordNotFoundErrorCode, "record not found")
)
