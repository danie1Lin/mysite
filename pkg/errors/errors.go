package errors

type Error interface {
	error
	Code() string
	Details() map[string][]string
	AddDetail(fieldName, message string)
}

type Err struct {
	subErr  error
	code    string
	message string
	details map[string][]string
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) Code() string {
	return e.code
}

func (e *Err) AddDetail(fieldName, message string) {
	msgArr := e.details[fieldName]
	if msgArr == nil {
		e.details[fieldName] = []string{message}
	} else {
		e.details[fieldName] = append(e.details[fieldName], message)
	}
}

func (e *Err) Details() map[string][]string {
	return e.details
}

func New(code string, message string) *Err {
	return &Err{
		code:    code,
		message: message,
		details: make(map[string][]string),
	}
}

func Code(err error) string {
	e, ok := err.(Error)
	if ok {
		return e.Code()
	}
	return ""
}

func Details(err error) map[string][]string {
	e, ok := err.(Error)
	if ok {
		return e.Details()
	}
	return nil
}

func AddDetail(err error, fieldName, message string) error {
	e, ok := err.(Error)
	if ok {
		e.AddDetail(fieldName, message)
		return nil
	}

	return &Err{
		subErr: err,
		details: map[string][]string{
			fieldName: []string{message},
		},
	}
}
