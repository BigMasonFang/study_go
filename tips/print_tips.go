package tips

import (
	"fmt"
	"study_go/tips/composite"
)

func init() {
	tips := []string{
		"struct_composite", "array_slice_composite",
		"map_composite",
	}
	fmt.Println("u can select from these tips")
	for i, v := range tips {
		fmt.Println("input ", i, "or", v)
	}
}

func PrintTips() {
	var input string
	fmt.Scanln(&input)

	switch input {
	case "0", "struct_composite":
		composite.PrintStructComposite()
	case "1", "array_slice_composite":
		composite.PrintArraySliceComposite()
	case "2", "map_composite":
		composite.PrintMapComposite()
	}
}
