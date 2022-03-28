package util

import (
	"github.com/pkg/errors"
	"os"
)

var ErrTest = errors.New("it is test")

func IsPathExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
