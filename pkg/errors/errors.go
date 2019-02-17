package errors

const (
	JSON_ERROR      = "json_error"
	REQ_ERROR       = "request_error"
	PARSE_ERROR     = "parse_error"
	DB_ERROR        = "json_error"
	NO_ERROR        = "no_error"
	UNDEFINED_ERROR = "undefined_error"
)

type Error struct {
	error
	Type string
}

func NewError(t string, err error) Error {
	return Error{err, t}
}

func ParseError(err error) string {
	errType := NO_ERROR
	if err != nil {
		if v, ok := err.(Error); !ok {
			errType = UNDEFINED_ERROR
		} else {
			errType = v.Type
		}
	}
	return errType
}
