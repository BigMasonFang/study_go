package study_go

import "fmt"

func PrintResult() {
	var name = "fangxiang"
	fmt.Println("name is", name)

	cnName := "方向"
	fmt.Println("name is", cnName)

	var (
		nickname string
		age      int
	)
	fmt.Println(nickname, age)

	cnName, enName, height := "方向", "Dime", 169
	fmt.Println(cnName, enName, height)
}
