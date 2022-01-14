package werror

import (
	"errors"
)

func Wrap(err error, cause error) ErrorNode {
	return &errorNode{
		error: err,
		cause: cause,
	}
}

func (e *errorNode) Wrap(err error) ErrorNode {
	return Wrap(err, e)
}

func (e *errorNode) Error() (res string) {
	_, ok := e.error.(*errorNode)
	if ok {
		res += e.error.Error()
	} else {
		res += e.error.Error() + "\n"
		c := errors.Unwrap(e.error)
		for c != nil {
			res += c.Error() + "\n"
			c = errors.Unwrap(c)
		}
	}

	c := e.Unwrap()
	for c != nil {
		res += c.Error() + "\n"
		c = errors.Unwrap(c)
	}

	return
}



func (e *errorNode) Unwrap() error {
	return e.cause
}

func (e *errorNode) Is(err error) bool {
	if errors.Is(e.error, err) {
		return true
	}

	c := e.Unwrap()
	for c != nil {
		if errors.Is(c, err) {
			return true
		}
		c = errors.Unwrap(c)
	}

	return false
}