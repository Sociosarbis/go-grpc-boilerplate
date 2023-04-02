package errgo

type wrapError struct {
	err error
	msg string
}

func Wrap(err error, msg string) error {
	if err == nil {
		return err
	}
	return &wrapError{
		err,
		msg,
	}
}

func (e *wrapError) Error() string {
	return e.msg + ": " + e.err.Error()
}

func (e *wrapError) Unwrap() error {
	return e.err
}
