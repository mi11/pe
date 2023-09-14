package main

import (
	"github.com/mi11/pe/pkg/config"
	"github.com/mi11/pe/pkg/srv"
)

func main() {
	env := config.LoadEnv()

	s, err := srv.NewServer(env)
	if err != nil {
		panic(err)
	}

	err = s.Start()
	if err != nil {
		panic(err)
	}
}
