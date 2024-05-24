package txdoc

import "fmt"

const (
	kProductionGateway = "https://docs.qq.com"
)

type Code int

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}

func (c Code) IsFailure() bool {
	return c != CodeSuccess
}

const (
	CodeSuccess Code = 0
)

type Error struct {
	Code    Code   `json:"ret"`
	Message string `json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func (e Error) IsSuccess() bool {
	return e.Code.IsSuccess()
}

func (e Error) IsFailure() bool {
	return e.Code.IsFailure()
}

type ErrorV3 struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (v3 ErrorV3) Error() Error {
	return Error{Code: v3.Code, Message: v3.Message}
}
