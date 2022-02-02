package erron

import "fmt"

var ErrNoTaskAvailable = fmt.Errorf("not task avaliable now, please try later")
