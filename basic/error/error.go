package error

import (
	"errors"
	"fmt"
)

var ErrBadInput = errors.New("bad input")

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput1(input string) error {
	if input == "error" {
		return fmt.Errorf("validateInput: %w", ErrBadInput)
	}
	return nil
}

func validateInput2(input string) error {
	if input == "error" {
		return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
	}
	return nil
}

func PrintErrorIs() {
	input := "error"
	err := validateInput1(input)
	if errors.Is(err, ErrBadInput) {
		fmt.Println("bad input error")
	}
}

func PrintErrorAs() {
	input := "error"
	err := validateInput2(input)
	var badInputError *BadInputError
	if errors.As(err, &badInputError) {
		fmt.Printf("bad input err occured : %s\n", badInputError)
	}

}
