package app

import (
	"gin-blog/pkg/logging"

	"github.com/astaxie/beego/validation"
)

func MakeErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
