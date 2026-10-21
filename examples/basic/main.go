package main

import (
	"fmt"
	lib "github.com/1Pas2/pas2-go-port-ranges"
)

func main() {
	ports, err := lib.Parse("80,443,8000-8003")
	fmt.Println(ports, err)
}
