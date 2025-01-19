package main

import (
	"fmt"

	"github.com/hakkiir/gator/internal/config"
)

func main() {
	cfg := config.Read()
	fmt.Println(cfg)
	config.SetUser("Iiro", cfg)
	cfg = config.Read()
	fmt.Println(cfg)
}
