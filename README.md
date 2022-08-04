# werror
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Frebirthlee%2Fwerror%2Fbadge%3Fref%3Dmain&style=flat)](https://actions-badge.atrox.dev/rebirthlee/werror/goto?ref=main)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/rebirthlee/werror)](https://github.com/RebirthLee/werror/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rebirthlee/werror)](https://go.dev/doc/go1.17)
[![license](https://img.shields.io/badge/license-BEER--WARE-green)](/LICENSE.md)

wrap error

## Install
`go get github.com/rebirthlee/werror`

## Example

### Simple

```go
package main

import (
	"errors"
	"fmt"
	"github.com/rebirthlee/werror"
)

var (
	errNotFound = errors.New("not found")
	errComputeFailed = errors.New("compute failed")
)

func getOne() (int64, error) {
	// get success
	// return 1, nil
	
	// get failed
	return 0, errNotFound
}

func compute() error {
	res, err := getOne()
	if err != nil {
		return werror.Wrap(errComputeFailed, err)
	}
	
	res++ // compute dummy
	return nil
}

func main() {
	err := compute()
	
	// is
	fmt.Println(errors.Is(err, errNotFound)) // true
	fmt.Println(errors.Is(err, errComputeFailed)) // true
}

```

### Advanced

```go
package main

import (
	"errors"
	"fmt"
	"github.com/rebirthlee/werror"
)

type myNotFoundError struct { }

func (*myNotFoundError) Error() string {
	return "not found"
}

type myComputeError struct { }

func (*myComputeError) Error() string {
	return "compute failed"
}


var (
	errNotFound error = &myNotFoundError{}
	errComputeFailed error = &myComputeError{}
)

func getOne() (int64, error) {
	// get success
	// return 1, nil
	
	// get failed
	return 0, errNotFound
}

func compute() error {
	res, err := getOne()
	if err != nil {
		return werror.Wrap(errComputeFailed, err)
	}
	
	res++ // compute dummy
	return nil
}

func main() {
	err := compute()
	
	// is
	fmt.Println(errors.Is(err, errNotFound)) // true
	fmt.Println(errors.Is(err, errComputeFailed)) // true
	
	// get error
	fmt.Println(errors.Unwrap(err)) // "compute failed"
	
	// get cause
	fmt.Println(werror.Cause(err)) // "not found"
	
	// as
	var node *werror.ErrorNode
	fmt.Println(errors.As(err, &node)) // true
	
	var myError1 *myNotFoundError
	fmt.Println(errors.As(err, &myError1)) // true

	var myError2 *myComputeError
	fmt.Println(errors.As(err, &myError2)) // true
}

```

## License
[`THE BEER-WARE LICENSE (Revision 42)`](http://en.wikipedia.org/wiki/Beerware)
