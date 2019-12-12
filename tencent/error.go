package tencent

import (
	"fmt"
	"regexp"
	"strconv"
)

type HttpRequestError struct {
	StatusCode int
}

func (e *HttpRequestError) Error() string {
	return fmt.Sprintf("failed to send http request, status code: %d",
		e.StatusCode)
}

type Error struct {
	*Response
}

func (e *Error) Error() string {
	return e.Message
}

var moduleCodeExp = regexp.MustCompile("\\((\\d+)\\).+")

func (e *Error) ModuleCode() int {
	if e.Message == "" {
		return 0
	}
	pieces := moduleCodeExp.FindStringSubmatch(e.Message)
	if pieces != nil && len(pieces) > 1 {
		code, err := strconv.Atoi(pieces[1])
		if err != nil {
			return 0
		}
		return code
	}
	return 0
}

type ConnectionError struct {
	cause string
}

func (e *ConnectionError) Error() string {
	return e.cause
}
