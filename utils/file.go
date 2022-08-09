package utils

import (
	"github.com/pkg/errors"
	"io/ioutil"
)

func ReadFile(path string) string {
	dataByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(errors.WithMessage(err, "打开文件失败"))
	}
	data := string(dataByte)

	return data
}
