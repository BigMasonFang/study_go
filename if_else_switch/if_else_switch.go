package study_go

import (
	"fmt"
)

type level string
type gradeInterval string

func grader(grade int) level {
	if grade < 60 {
		return "F"
	} else if grade < 70 {
		return "D"
	} else if grade < 80 {
		return "C"
	} else if grade < 90 {
		return "B"
	} else if grade <= 100 {
		return "A"
	}
	return "error input"
}

func passOrFail(grade int) level {
	switch {
	case grade < 60:
		{
			return "F"
		}
	default:
		{
			return "P"
		}
	}
}

func gradeIntervals(level level) gradeInterval {
	switch level {
	case "A":
		return "[90, 100]"
	case "B":
		return "[80, 90)"
	case "C":
		return "[70, 80)"
	case "D":
		return "[60, 70)"
	case "F":
		return "[0, 60]"
	}
	return "error input"
}

func PrintLogicControl() {
	english, math, music := 80, 95, 50
	fmt.Println("english grade is", grader(english))
	fmt.Println("math grade is", grader(math))
	fmt.Println("music grade is", grader(music))
	fmt.Printf("music pass or fail? %v\n", passOrFail(music))
	fmt.Printf("grader interval is %v if level is C\n", gradeIntervals("C"))
}
