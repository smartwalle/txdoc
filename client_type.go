package txdoc

import (
	"encoding/json"
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
	Code    Code   `json:"code"`
	Message string `json:"message"`
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

func (e *Error) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Ret     Code   `json:"ret"`
		Msg     string `json:"msg"`
		Code    Code   `json:"code"`
		Message string `json:"message"`
	}{}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.Ret > 0 {
		e.Code = aux.Ret
		e.Message = aux.Msg
	} else if aux.Code > 0 {
		e.Code = aux.Code
		e.Message = aux.Message
	}
	return nil
}
