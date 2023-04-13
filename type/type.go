package Type

import (
	"github.com/davecgh/go-spew/spew"
)

func PrintType() {
	var a int = 8
	var b int8 = 2
	spew.Dump(a + int(b))
}
