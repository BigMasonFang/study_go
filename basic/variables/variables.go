package variables

import (
	"github.com/davecgh/go-spew/spew"
)

func PrintVars() {
	const doubler int = 2
	var a int = 9 // standard
	var b int     // declare fisrt, will assign default
	b = 8         // assign value later
	var c = 10    // type infer

	var d, e, f = "2", 'x', 7 // group create with infer
	var (
		x int = 8
		y     = 7
	)
	spew.Dump(a, b, c, d, e, f, x, y)

	speed := x * y // candy
	old_speed := speed
	speed = doubler * speed
	// speed := 8 // could not create same var multiple times in one context
	spew.Dump(old_speed, speed) // var create is not reference
	spew.Printf("old speed is %v, new speed is %v\n", old_speed, speed)

	var g string         // default ""
	var h int            // default 0
	var slice1 []int     // nil
	var array1 [3]int    // nil
	var array2 [3]string // [3]string of ""
	spew.Dump(g, h, slice1, array1, array2)

}
