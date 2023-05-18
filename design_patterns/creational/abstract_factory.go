package creational

import "fmt"

// abstract product
type Button interface {
	Click()
}

type CheckBox interface {
	Check()
}

// absstract factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

// concrete product
type WinButton struct{}
type MacButton struct{}
type WinCheckBox struct{}
type MacCheckBox struct{}

func (w *WinButton) Click() {
	fmt.Println("win button clicked")
}

func (w *MacButton) Click() {
	fmt.Println("mac button clicked")
}

func (w *WinCheckBox) Check() {
	fmt.Println("win checkbox checked")
}

func (w *MacCheckBox) Check() {
	fmt.Println("mac checkbox checked")
}

// concrete factory
type WinFactory struct{}
type MacFactory struct{}

func (w *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (w *WinFactory) CreateCheckBox() CheckBox {
	return &WinCheckBox{}
}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (m *MacFactory) CreateCheckBox() CheckBox {
	return &MacCheckBox{}
}

func PrintAbstractFactory() {
	var factory GUIFactory = &WinFactory{}

	button := factory.CreateButton()
	checkBox := factory.CreateCheckBox()

	button.Click()
	checkBox.Check()
}
