package error

type CustomError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func New(name string) CustomError {
	return CustomError{name, ""} // enforce the default value here
}

func NewCode(name string, code string) CustomError {
	return CustomError{name, code} // enforce the default value here
}
