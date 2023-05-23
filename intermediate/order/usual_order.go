package Order

import "fmt"

func ff() int {
	fmt.Println("calling ff")
	return 1
}

func gg(a, b, c int) int {
	fmt.Println("calling gg")
	return 2
}
func hh() int {
	fmt.Println("calling hh")
	return 3
}
func ii() int {
	fmt.Println("calling ii")
	return 1
}
func jj() int {
	fmt.Println("calling jj")
	return 1
}
func kk() bool {
	fmt.Println("calling kk")
	return true
}
func PrintUsualOrder() {
	var y = []int{11, 12, 13}
	var x = []int{21, 22, 23}

	var c chan int = make(chan int)
	go func() {
		c <- 1
		fmt.Println("sended")
	}()

	y[ff()], _ = gg(hh(), ii()+x[jj()], <-c), kk()
	// ff, hh, ii, jj, sended, gg, kk
}
