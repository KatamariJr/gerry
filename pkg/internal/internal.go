package internal

// Error (the internal.Error type) is used to differentiate a usual execution error from one that contains internal
// data which shouldnt leak to public endpoints. This has little extra functionality besides a typical error, but is
// useful for checking via errors.Is() to see if an error contains an internal message.
type Error struct {
	internalErr error
}

func NewInternalError(err error) Error {
	return Error{
		internalErr: err,
	}
}

func (e *Error) Error() string {
	return e.internalErr.Error()
}
