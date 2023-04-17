package method

import "fmt"

type Cordinate struct {
	x, y, z float64
}

// with * u can modify the struct instance (object)

// Methods with pointer receivers can modify the value to which the receiver points.
//  Since methods often need to modify their receiver,
//  pointer receivers are more common than value receivers.

// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call.
// This can be more efficient if the receiver is a large struct
func (cor *Cordinate) ShiftBy(x, y, z float64) {
	cor.x += x
	cor.y += y
	cor.z += z
}

// func without pointer receviers can also do the same
// but notice if want to modify struct, also need pass pointer
func shiftByFunc(cor *Cordinate, x, y, z float64) {
	cor.x += x
	cor.y += y
	cor.z += z
}

func (cor *Cordinate) Reset() {
	cor.x, cor.y, cor.z = 0, 0, 0
}

func InitCordinate() *Cordinate {
	return &Cordinate{0, 0, 0}
}

func PrintMethod() {
	cord1 := Cordinate{1, 2, 3}
	cord1P := &cord1
	cord2 := cord1

	// methods with pointer receivers take either a value or a pointer as the receiver when they are called
	cord1.ShiftBy(1, 2, 3)
	fmt.Println(cord1)
	cord1P.ShiftBy(1, 2, 3)
	fmt.Println(cord1)

	// functions with a pointer argument must take a pointer
	shiftByFunc(&cord2, 1, 2, 3)
	fmt.Println(cord2)
	// will error if
	// shiftByFunc(cord2,1,2,3)

}
