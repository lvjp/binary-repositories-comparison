package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed data/*.yaml
var dataFiles embed.FS

func main() {
	fmt.Println("Pouet !!!")
}
