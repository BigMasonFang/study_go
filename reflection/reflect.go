package reflection

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func PrintReflect() {
	user := User{"Tom", 17}

	fmt.Println(reflect.TypeOf(user))
	fmt.Println(reflect.ValueOf(user))

	t := reflect.TypeOf(user)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s %s\n", f.Name, f.Type)
	}

	v := reflect.ValueOf(&user).Elem()
	v.FieldByName("Name").SetString("Jerry")
	v.FieldByName("Age").SetInt(9)
	fmt.Println(user)
}
