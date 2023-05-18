package structural

import (
	"errors"
	"fmt"
)

// Abastraction
type Device interface {
	IsEnable() bool
	Enable() error
	Disable() error
	GetChannel() int
	SetChannel(channel int) error
}

// bridge
type BasicDevice struct {
	Channel int
	Enabled bool
}

func (d *BasicDevice) IsEnable() bool {
	return d.Enabled
}

func (d *BasicDevice) Enable() error {
	d.Enabled = true
	return nil
}

func (d *BasicDevice) Disable() error {
	d.Enabled = false
	return nil
}

func (d *BasicDevice) GetChannel() int {
	return d.Channel
}

func (d *BasicDevice) SetChannel(channel int) error {
	d.Channel = channel
	return nil
}

// Implementation
type TV struct {
	Device
}

type DVDPlayer struct {
	Device
}

func (tv *TV) TurnOn() error {
	if !tv.IsEnable() {
		err := tv.Enable()
		if err != nil {
			return err
		}
	}
	return nil
}

func (tv *TV) TurnOff() error {
	if tv.IsEnable() {
		err := tv.Disable()
		if err != nil {
			return err
		}
	}
	return nil
}

func (tv *TV) ChangeChannel(c int) error {
	if tv.IsEnable() {
		err := tv.SetChannel(c)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (dvd *DVDPlayer) TurnOn() error {
	if !dvd.IsEnable() {
		err := dvd.Enable()
		if err != nil {
			return err
		}
	}
	return nil
}

func (dvd *DVDPlayer) TurnOff() error {
	if dvd.IsEnable() {
		err := dvd.Disable()
		if err != nil {
			return err
		}
	}
	return nil
}

func (dvd *DVDPlayer) ChangeChannel(c int) error {
	return errors.New("DVD player can not change channel")
}

func PrintBridgeExec() {
	tv := &TV{&BasicDevice{}}
	dvd := &DVDPlayer{&BasicDevice{}}

	fmt.Println("tv's status", tv.IsEnable())
	fmt.Println("turn on tv")
	tv.TurnOn()
	if tv.IsEnable() {
		fmt.Println("tv is on")
	} else {
		fmt.Println("tv is off")
	}
	fmt.Println("tv's channel is", tv.GetChannel())
	fmt.Println("change channel to 5")
	tv.ChangeChannel(5)
	fmt.Println("tv's channel is", tv.GetChannel())

	fmt.Println("dvd's status", dvd.IsEnable())
	fmt.Println("turn on dvd")
	dvd.TurnOn()
	if dvd.IsEnable() {
		fmt.Println("dvd is on")
	} else {
		fmt.Println("dvd is off")
	}
	fmt.Println("dvd's channel is", tv.GetChannel())
	fmt.Println("change channel to 5")
	err := dvd.ChangeChannel(5)
	if err != nil {
		fmt.Println("change dvd's channel meet error", err)
	}
}
