package endpoints

type Response interface {
	Code() string
	Error() string
	Meta() map[string]string
	Message() string
	StatusCode() int
}

type response struct {
	statusCode int
	code       string
	error      error
	meta       map[string]string
	message    string
}

func (r *response) StatusCode() int {
	return r.statusCode
}

func (r *response) Code() string {
	return r.code
}

func (r *response) Error() string {
	if r.error == nil {
		return ""
	}
	return r.error.Error()
}

func (r *response) Meta() map[string]string {
	return r.meta
}

func (r *response) Message() string {
	return r.message
}
