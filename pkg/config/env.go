package config

import "os"

type Env struct {
	Port string
	Name string
}

func LoadEnv() *Env {
	config := &Env{
		Port: "5000",
		Name: "User",
	}

	if value, exists := os.LookupEnv("PORT"); exists {
		config.Port = value
	}

	if value, exists := os.LookupEnv("NAME"); exists {
		config.Name = value
	}

	return config
}
