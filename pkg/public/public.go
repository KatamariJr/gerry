package public

// PublicError is an error type used to specify an error that should be served to users while obfuscating an actual error message, possibly for logging. Fo example, returning the
// actual Error() result in a 500 response from an API, and using the InternalError() function result for logging to an error tracking system.

type Error struct {
	err                 error
	PublicFacingMessage string
}

func NewPublicError(err error, publicMessage string) Error {
	return Error{
		err:                 err,
		PublicFacingMessage: publicMessage,
	}
}

func (e *Error) Error() string {
	return e.PublicFacingMessage
}

func (e *Error) InternalError() error {
	return e.err
}
