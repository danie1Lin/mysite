package endpoints

import "mysite/pkg/errors"

const (
	DishandlableErrorCode       string = "dishandlable_error"
	FieldNotAllowEmptyErrorCode string = "field_not_allow_empty_error"
)

var (
	DishandlableError       error = errors.New(DishandlableErrorCode, "dishandlable error")
	FieldNotAllowEmptyError error = errors.New(FieldNotAllowEmptyErrorCode, "field not allow empty error")
)
