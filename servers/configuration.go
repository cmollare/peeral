package main

// Configuration for network connection
type Configuration struct {
	Port int
}

// NewConfiguration create configuration object for the app
func NewConfiguration() Configuration {
	return Configuration{Port: 8000}
}
