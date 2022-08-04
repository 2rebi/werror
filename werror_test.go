package werror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestError1(text string) error {
	return &testError1{text}
}

type testError1 struct {
	s string
}

func (e *testError1) Error() string {
	return e.s
}

func newTestError2(text string) error {
	return &testError2{text}
}

type testError2 struct {
	s string
}

func (e *testError2) Error() string {
	return e.s
}

func TestWrap(t *testing.T) {
	{
		err := newTestError1("test error")
		assert.Equal(t, Wrap(err, nil), err)
		assert.Equal(t, Wrap(nil, err), err)
	}

	{
		assert.Nil(t, Wrap(nil, nil))
	}

	{
		err1, err2 := newTestError1("test error 1"), newTestError1("test error 2")
		werr := Wrap(err1, err2)

		var wrapper *ErrorNode
		assert.True(t, errors.As(werr, &wrapper))

		assert.Equal(t, wrapper.error, err1)
		assert.Equal(t, wrapper.cause, err2)
	}
}

func TestCause(t *testing.T) {
	{
		err1, err2 := newTestError1("test error 1"), newTestError1("test error 2")
		werr := Wrap(err1, err2)

		var wrapper *ErrorNode
		assert.True(t, errors.As(werr, &wrapper))

		assert.Equal(t, Cause(werr), wrapper.cause)
		assert.Equal(t, Cause(werr), err2)
		assert.Nil(t, Cause(err1))
		assert.Nil(t, Cause(err2))
	}
}

func TestErrorNode_Error(t *testing.T) {
	/*
	   root (ErrorNode)
	   	error: _node__node_err1_err2__err3_ (ErrorNode)
	   				error: _node_err1_err2_ (ErrorNode)
	   							error: error1
	   							cause: error2
	   				cause: error3
	   	cause: _node_err4_err5_ (ErrorNode)
	   				error: error4
	   				cause: error5
	*/
	var (
		error1 = newTestError1("error1")
		error2 = newTestError1("error2")
		error3 = newTestError1("error3")
		error4 = newTestError1("error4")
		error5 = newTestError1("error5")

		_node_err1_err2_ = Wrap(
			error1,
			error2,
		)

		_node__node_err1_err2__err3_ = Wrap(
			_node_err1_err2_,
			error3,
		)

		_node_err4_err5_ = Wrap(
			error4,
			error5,
		)

		root = Wrap(
			_node__node_err1_err2__err3_,
			_node_err4_err5_,
		)
	)

	assert.Equal(t, root.Error(), "error1\nerror2\nerror3\nerror4\nerror5")
}

func TestErrorNode_Unwrap(t *testing.T) {
	{
		err1, err2 := newTestError1("test error 1"), newTestError1("test error 2")
		werr := Wrap(err1, err2)

		var wrapper *ErrorNode
		assert.True(t, errors.As(werr, &wrapper))

		assert.Equal(t, errors.Unwrap(werr), wrapper.error)
		assert.Equal(t, errors.Unwrap(werr), err1)
		assert.Equal(t, errors.Unwrap(werr), wrapper.Unwrap())
	}
}

func TestErrorNode_Cause(t *testing.T) {
	{
		err1, err2 := newTestError1("test error 1"), newTestError1("test error 2")
		werr := Wrap(err1, err2)

		var wrapper *ErrorNode
		assert.True(t, errors.As(werr, &wrapper))

		assert.Equal(t, Cause(werr), wrapper.Cause())
	}
}

func TestErrorNode_Is(t *testing.T) {
	{
		err1, err2 := newTestError1("test error 1"), newTestError2("test error 2")
		werr := Wrap(err1, err2)

		assert.True(t, errors.Is(werr, err1))
		assert.True(t, errors.Is(werr, err2))
	}
}

func TestErrorNode_As(t *testing.T) {
	{
		err1, err2 := newTestError1("test error 1"), newTestError2("test error 2")
		werr := Wrap(err1, err2)

		var wrapper *ErrorNode
		assert.True(t, errors.As(werr, &wrapper))

		var t1 *testError1
		var t2 *testError2
		assert.True(t, errors.As(werr, &t1))
		assert.True(t, errors.As(werr, &t2))
	}
}
