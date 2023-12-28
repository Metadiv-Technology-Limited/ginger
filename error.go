package ginger

type IError interface {
	GetCode() string
	SetCode(string)
	GetMessage() string
	SetMessage(string)
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) GetCode() string {
	return e.Code
}

func (e *Error) SetCode(code string) {
	e.Code = code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}
