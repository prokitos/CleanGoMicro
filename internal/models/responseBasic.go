package models

import (
	"errors"
	"strconv"
)

type ResponseGlobal struct {
	Description string `json:"description"       		 example:"description"`
	Code        int    `json:"code"               		 example:"status"`
}

func (instance ResponseGlobal) BadRequest() error {
	instance.Code = 400
	instance.Description = "invalid sended data"
	return errors.New(instance.getResult())
}
func (instance ResponseGlobal) InternalError() error {
	instance.Code = 400
	instance.Description = "internal error"
	return errors.New(instance.getResult())
}
func (instance ResponseGlobal) getResult() string {
	result := "Code:" + strconv.Itoa(instance.Code) + "; Description:" + instance.Description
	return result
}
