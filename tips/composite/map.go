package composite

import "fmt"

type Cord struct {
	x, y float64
}

func PrintMapComposite() {
	cords := []Cord{
		// can negelect type inside slice
		{1.0, 2.0}, {3.0, 4.0},
	}

	locationCords := map[string]Cord{
		// can negelect type inside map
		"A": {1.0, 2.0},
		"B": {3.0, 4.0},
	}

	locationCordsP := map[string]*Cord{
		// can negelect type inside map
		"A": {1.0, 2.0}, // also the pointer
		"B": {3.0, 4.0},
	}

	fmt.Println(cords)
	fmt.Println(locationCords)
	fmt.Println(locationCordsP)
}
