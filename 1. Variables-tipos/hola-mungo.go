package main

import "fmt"

func main() {
	nb, err := fmt.Println("Hola mundo")
	_, _ = fmt.Println(nb, err)
	// nb2, err2 := fmt.Println(nb, err)
}