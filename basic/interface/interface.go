package _interface

import (
	"errors"
	"fmt"

	m "study_go/basic/method"
)

type Resetter interface {
	Reset()
}

type PlayerChecker interface {
	ValidateHealth() error
}

type InvalidHealthError struct {
	health int
}

func (err InvalidHealthError) Error() string {
	return fmt.Sprintf("bad health: %v", err.health)
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

func (p *Player) ValidateHealth() error {
	if p.health < 0 || p.health > 100 {
		return fmt.Errorf("invalid health %w", &InvalidHealthError{p.health})
	}
	return nil
}

func (p *Player) DeductHealth(v int) {
	p.health -= v
}

// func Reset can receive Resetter type as input param
// in another word, any type which realized Reset() method
// can use as the param, ie "Player" type
func Reset(r Resetter) {
	r.Reset()
}

func Validate(r PlayerChecker) error {
	err := r.ValidateHealth()
	return err
}

func PrintInterface() {
	player := Player{100, *m.InitCordinate()}
	fmt.Println(player)
	player.position.ShiftBy(1, 2, 3)
	player.DeductHealth(30)
	fmt.Println(player)
	Reset(&player)
	fmt.Println(player)
	player.DeductHealth(110)
	err := Validate(&player)
	var invalidHealthError *InvalidHealthError
	if err != nil {
		if errors.As(err, &invalidHealthError) {
			fmt.Printf("player's %v health is not valid", player.health)
		}
	}
}
