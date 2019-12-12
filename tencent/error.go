package tencent

import "fmt"

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

type ConnectionError struct {
	cause string
}

func (e *ConnectionError) Error() string {
	return e.cause
}
