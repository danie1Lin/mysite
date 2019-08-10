package cache

import "mysite/pkg/errors"

const (
	DataInvalidErrorCode         string = "data_invalid_error"
	SessionIDDuplicatedErrorCode        = "session_id_duplicated_error"
)

var (
	DataInvalidError         error = errors.New(DataInvalidErrorCode, "data invalid")
	SessionIDDuplicatedError       = errors.New(SessionIDDuplicatedErrorCode, "session id duplicated error")
)
