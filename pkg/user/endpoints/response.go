package endpoints

type Response interface {
	Code() string
	Error() string
	Meta() map[string]string
}

type response struct {
	code  string
	error error
	meta  map[string]string
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
