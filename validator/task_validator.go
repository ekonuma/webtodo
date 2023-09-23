package validator

import "github.com/ekonuma/webtodo/model"

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct {}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}