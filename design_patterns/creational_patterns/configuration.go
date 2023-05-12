package creational_patterns

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost   string
	DBPort   int
	DBName   string
	LogLevel string
}

type ComplexConfig struct {
	Host     string
	Port     int
	LogLevel string
}

type Option func(*ComplexConfig)

func WithHost(host string) Option {
	return func(cc *ComplexConfig) {
		cc.Host = host
	}
}

func WithPort(port int) Option {
	return func(cc *ComplexConfig) {
		cc.Port = port
	}
}

func WithLogLevel(logLevel string) Option {
	return func(cc *ComplexConfig) {
		cc.LogLevel = logLevel
	}
}

func LoadConfig() (*Config, error) {
	// get config from env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	logLevel := os.Getenv("LOG_LEVEL")

	// parse db port
	var port int
	if dbPort != "" {
		_, err := fmt.Scanf(dbPort, "%d", &port)
		if err != nil {
			return nil, err
		}
	}

	// return config instance
	return &Config{
		DBHost:   dbHost,
		DBPort:   port,
		DBName:   dbName,
		LogLevel: logLevel,
	}, nil
}

func NewConfig(opts ...Option) *ComplexConfig {
	c := &ComplexConfig{
		Host:     "localhost",
		Port:     8080,
		LogLevel: "INFO",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func PrintComplexConfig() {
	c1 := NewConfig()
	fmt.Printf("Default config: %v\n", c1)

	c2 := NewConfig(WithHost("127.0.0.1"), WithPort(80), WithLogLevel("DEBUG"))
	fmt.Printf("Custom config: %v\n", c2)
}
