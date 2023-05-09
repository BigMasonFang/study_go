package test

import (
	"regexp"
)

func IsValidEmail(addr string) bool {
	re, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to complie regrex")
	} else {
		return re.Match([]byte(addr))
	}
}
