package main

import (
	"backend/config"
	"fmt"
)

func main() {
	c, err := config.Load()
	fmt.Println(c, err)
}
