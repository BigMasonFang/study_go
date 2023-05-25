package funcs

import "fmt"

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	// using closure
	// Closure is a function that references variables declared outside its defining scope.
	//  Inside the function, you can access external variables,
	// even when the function is called outside, the external variables keep their values.
	//  The closure "captures" the state of the external variables, and can modify these variables inside the function,
	return func(y int) int {
		return times(x, y)
	}
}

// 1 curried function
func PrintCurried() {
	timesTwo := partialTimes(2)
	timesThree := partialTimes(3)

	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(5))
}

// 2 functor
type IntSliceFunctor interface {
	Fmap(fn func(int) int) IntSliceFunctor
}

type IntSliceFunctorImpl struct {
	ints []int
}

func (isf IntSliceFunctorImpl) Fmap(fn func(int) int) IntSliceFunctor {
	newInts := make([]int, len(isf.ints))
	for i, e := range isf.ints {
		retInt := fn(e)
		newInts[i] = retInt
	}
	return IntSliceFunctorImpl{newInts}
}

func NewIntSliceFunctor(slice []int) IntSliceFunctor {
	return IntSliceFunctorImpl{slice}
}

func PrintFunctor() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("init a functor from int slice %v\n", s)
	f := NewIntSliceFunctor(s)
	fmt.Printf("original functor: %v\n", f)

	mapperSquare := func(i int) int { return i * i }
	mappedSquare := f.Fmap(mapperSquare)
	fmt.Printf("mapped square functor: %v\n", mappedSquare)

	mapperCube := func(i int) int { return i * i * i }
	mappedCube := f.Fmap(mapperCube)
	fmt.Printf("mapped cube functor: %v\n", mappedCube)

	fmt.Printf("original functor: %v\n", f)
	fmt.Printf("composite functor: %v\n", f.Fmap(mapperSquare).Fmap(mapperCube))
}

// 3 recursion and Continuation-passing style CPS(no return)
// simple example
func Max(m, n int) int {
	if n > m {
		return n
	} else {
		return m
	}
}

func MaxCPS(m, n int, f func(int)) {
	if n > m {
		f(n)
	} else {
		f(m)
	}
}

func PrintRecursionSimple() {
	MaxCPS(10, 6, func(n int) { fmt.Println(n) })
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

func FactorialCPS(n int, f func(n int)) {
	if n == 1 {
		f(1)
	} else {
		fmt.Printf("%d * %T\n", n, f)
		FactorialCPS(n-1, func(z int) { f(n * z) })
	}
}

func PrintRecursionCPS() {
	FactorialCPS(5, func(z int) { fmt.Printf("%d\n", z) })
}
