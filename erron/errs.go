package erron

import (
	"github.com/pkg/errors"
)

var ErrNoTaskAvailable = errors.Errorf("not task avaliable now, please try later")
