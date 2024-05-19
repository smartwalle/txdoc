package txdoc

import (
	"fmt"
)

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
	Ret Code   `json:"ret"`
	Msg string `json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Ret, e.Msg)
}

func (e Error) IsSuccess() bool {
	return e.Ret.IsSuccess()
}

func (e Error) IsFailure() bool {
	return e.Ret.IsFailure()
}
