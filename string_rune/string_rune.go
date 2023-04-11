package study_go

import "github.com/davecgh/go-spew/spew"

func PrintStringRune() {
	const nihong = "アメリカ合衆国ドルAmerica中国"
	disprintable := `😂`

	for index, runeV := range nihong {
		spew.Printf("%U %v starts at byte position %d\n", runeV, runeV, index)
	}

	// not working if using for loop
	// spew.Println(len(nihong))
	// for i := 0; i < len(nihong); i++ {
	// 	spew.Printf("%U %v\n", string(nihong[i]))
	// }

	for _, runeV := range disprintable {
		spew.Printf("%U %v \n", runeV, runeV)
	}

	spew.Dump(disprintable)
	spew.Println(`let's print in "Golang"`) // raw literal
}
