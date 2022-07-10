package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.yaml
var rawData []byte

func main() {
	fmt.Println("Pouet !!!")
	fmt.Println(string(rawData))
}
