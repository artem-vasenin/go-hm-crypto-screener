package main

import (
	"fmt"
	"hm-screener/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// test the environment
	fmt.Print(cfg)
}
