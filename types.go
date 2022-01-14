package werror

type ErrorNode interface {
	error
	Wrap(cause error) ErrorNode
}

type errorNode struct {
	error // current
	cause error // next
}
