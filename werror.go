package werror

import (
	"errors"
	"strings"
)

func Wrap(err error, cause error) error {
	flag := byte(0)
	if err != nil {
		flag |= 0x01
	}

	if cause != nil {
		flag |= 0x02
	}

	switch flag {
	case 0x01:
		return err
	case 0x02:
		return cause
	case 0x03:
		return &ErrorNode{
			error: err,
			cause: cause,
		}
	}

	return nil
}

func Cause(err error) error {
	if c, ok := err.(interface{ Cause() error }); ok {
		return c.Cause()
	}

	return nil
}

type ErrorNode struct {
	error error // current
	cause error // next
}

func (e *ErrorNode) Error() string {
	var b strings.Builder
	e.walk(&b)
	return b.String()
}

func (e *ErrorNode) walk(b *strings.Builder) {
	walkNode(b, e.error)
	b.WriteRune('\n')
	walkNode(b, e.cause)
}

func walkNode(b *strings.Builder, err error) {
	if n, ok := err.(*ErrorNode); ok {
		n.walk(b)
	} else {
		b.WriteString(err.Error())
	}
}

func (e *ErrorNode) Unwrap() error {
	return e.error
}

func (e *ErrorNode) Is(err error) bool {
	return errors.Is(e.error, err) || errors.Is(e.cause, err)
}

func (e *ErrorNode) As(dst interface{}) bool {
	return errors.As(e.error, dst) || errors.As(e.cause, dst)
}

func (e *ErrorNode) Cause() error {
	return e.cause
}
