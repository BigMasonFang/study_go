package control

import (
	"fmt"
	"math/rand"
)

type UserDefineErr struct {
	input string
}

func (e *UserDefineErr) Error() string {
	return "user define err"
}

func DoSomethingHappy() (string, error) {
	// good practice

	condition1 := "good"
	// some success logic
	// 1. sucess logic will not in if else control

	if condition1 != "good" {
		// condition 1 error logic
		// 2. find err, return immediately
		return "", fmt.Errorf("%w", &UserDefineErr{"condition 1"})
	}

	// some success logic
	// 3. sucess logic is left indented

	condition2 := rand.Intn(4)
	if condition2 == 0 {
		// condition 2 error logic 2
		return "", fmt.Errorf("%w", &UserDefineErr{"condition 2"})
	}
	// some success logic
	// 4. happy path's return is usually in the last line of a func
	return "success", nil
}

func DoSomethingSad() (string, error) {
	// bad practice
	condition1 := "good"
	if condition1 == "good" {
		// some success logic

		condition2 := rand.Intn(4)
		if condition2 != 0 {
			// some success logic
			return "success", nil
		} else {
			// condition 2 error logic 2
			return "", fmt.Errorf("%w", &UserDefineErr{"condition 2"})
		}
	} else {
		// condition 1 error logic
		return "", fmt.Errorf("%w", &UserDefineErr{"condition 1"})
	}
}
