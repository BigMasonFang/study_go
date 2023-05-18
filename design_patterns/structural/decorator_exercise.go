package structural

import "fmt"

type ObjectInterface interface {
	Do() error
}

type ConcreteObj struct{}

func (cO ConcreteObj) Do() error {
	defer fmt.Println("at least did it!!")
	// implementation
	fmt.Println("ok")
	return nil
}

// abstract
type DecoratorInterface interface {
	Do() error
}

type SimpleDecorator struct {
	Object ObjectInterface
}

func (simpleD SimpleDecorator) Do() error {
	// added func go's there
	fmt.Println("add some easy print job like logging")

	return simpleD.Object.Do()
}

type ExampleLogger struct {
	Level        string
	IntervalTime int
}

func (eLogger *ExampleLogger) Run() {
	fmt.Printf("start at %s, interval time is %v\n", eLogger.Level, eLogger.IntervalTime)
}

func (eLogger *ExampleLogger) SetLevel(level string) {
	eLogger.Level = level
}

func (eLogger *ExampleLogger) SetCronInterval(interval int) {
	eLogger.IntervalTime = interval
}

type ComplexDecorator struct {
	Object ObjectInterface
	Logger *ExampleLogger
	Level  string
}

func (complexD ComplexDecorator) Do() error {
	defer complexD.Logger.Run()
	fmt.Println("the log level is", complexD.Level)

	switch complexD.Level {
	case "DEBUG", "TEST":
		fmt.Println("the logger will log at debug level and will affect performance")
		complexD.Logger.SetLevel("DEBUG")
		complexD.Logger.SetCronInterval(3600)
	case "INFO":
		complexD.Logger.SetLevel("INFO")
		complexD.Logger.SetCronInterval(3600)
	case "WARNING":
		complexD.Logger.SetLevel("WARNING")
		complexD.Logger.SetCronInterval(3600 * 12)
	case "PROD", "ERROR":
		complexD.Logger.SetLevel("ERROR")
		complexD.Logger.SetCronInterval(3600 * 24)
	}

	return complexD.Object.Do()
}

func PrintDecoratorExercise() {
	obj := &ConcreteObj{}
	simpleDecoratedObj := SimpleDecorator{obj}
	complexDecoratorObj := ComplexDecorator{obj, &ExampleLogger{}, "TEST"}

	simpleDecoratedObj.Do()
	complexDecoratorObj.Do()
}
