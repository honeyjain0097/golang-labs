package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
)

type config struct {
	Foo string `env:"FOO"`
	Bar string `env:"BAR"`
}

func main() {
	fmt.Println("FOO:", os.Getenv("FOO")) // this is native example
	fmt.Println("BAR:", os.Getenv("BAR"))

	// ------------------------------------------------ // this is recommended example
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	// or you can use generics
	cfg, err := env.ParseAs[config]()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}
