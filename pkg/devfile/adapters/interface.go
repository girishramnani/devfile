package adapters

import (
	"io"

	"github.com/girishramnani/devfile/pkg/devfile/adapters/common"
)

type PlatformAdapter interface {
	Push(parameters common.PushParameters) error
	DoesComponentExist(cmpName string) bool
	Delete(labels map[string]string) error
	Log(follow, debug bool) (io.ReadCloser, error)
	Exec(command []string) error
}
