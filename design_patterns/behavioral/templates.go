package behavioral

import "fmt"

// abstract game interface
type Game interface {
	Init()
	Start()
	End()
}

// a base game (template) to implement template
type GameBase struct {
	Game
}

func (g *GameBase) Play() {
	g.Init()
	g.Start()
	g.End()
}

// concrete Game basketball
type BasketballGame struct {
	GameBase
}

func (bg *BasketballGame) Init() {
	fmt.Println("init basketball game")
}

func (bg *BasketballGame) Start() {
	fmt.Println("start basketball game")
}

func (bg *BasketballGame) End() {
	fmt.Println("End basketball game")
}

// concrete Game basketball
type FootballGame struct {
	GameBase
}

func (fg *FootballGame) Init() {
	fmt.Println("init football game")
}

func (fg *FootballGame) Start() {
	fmt.Println("start football game")
}

func (fg *FootballGame) End() {
	fmt.Println("End football game")
}

func PrintTemplates() {
	basketballGame := &BasketballGame{}
	basketballGame.Game = basketballGame
	basketballGame.GameBase.Play()

	fmt.Println()

	footballGame := &FootballGame{}
	footballGame.Game = footballGame
	footballGame.GameBase.Play()
}
