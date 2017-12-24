package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "avatar.jpg"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Arquivo não existe")
	} else {
		fmt.Println("Arquivo encontrado")
	}
}
