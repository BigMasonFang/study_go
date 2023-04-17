package _interface

import (
	"fmt"

	m "github.com/BigMasonFang/study_go/method"
)

type Resetter interface {
	Reset()
}

type Player struct {
	health   int
	position m.Cordinate
}

func (p *Player) Reset() {
	p.health = 100
	// not working because it is lower case
	// either change it to upper case which is exportable
	// or wirte init func
	p.position.Reset()
}

func (p *Player) DeductHealth(v int) {
	p.health -= v
}

// func Reset can receive param which is "Resetter" interface
// any type which satistied "Resetter" interface (in this case implement Reset() method)
// can use as the param, ie "Player" type
func Reset(r Resetter) {
	r.Reset()
}

func PrintInterface() {
	player := Player{100, *m.InitCordinate()}
	fmt.Println(player)
	player.position.ShiftBy(1, 2, 3)
	player.DeductHealth(30)
	fmt.Println(player)
	Reset(&player)
	fmt.Println(player)
}
