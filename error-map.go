package ginger

const defaultLocale = "default"

var errMapObj = errMap{}

type errMap map[string]map[string]string

func RegisterError(code, message string, locale ...string) {
	if len(locale) == 0 {
		locale = append(locale, defaultLocale)
	}
	if _, ok := errMapObj[locale[0]]; !ok {
		errMapObj[locale[0]] = make(map[string]string)
	}
	errMapObj[locale[0]][code] = message
}

func (e errMap) GetError(code string, locale ...string) IError {
	if len(locale) == 0 {
		locale = append(locale, defaultLocale)
	}
	if _, ok := errMapObj[locale[0]]; !ok {
		return &Error{
			Code: code,
		}
	}
	if _, ok := errMapObj[locale[0]][code]; !ok {
		return &Error{
			Code: code,
		}
	}
	return &Error{
		Code:    code,
		Message: errMapObj[locale[0]][code],
	}
}
