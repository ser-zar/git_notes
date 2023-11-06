package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Who are you?")
	fmt.Printf("> ")
	буфер := make([]byte, 120)
	длина, _ := os.Stdin.Read(буфер)
	Ω := длина
	ответ := string(буфер[:Ω-1])
	fmt.Printf("Какое длинное имя ... целых %d байт \n", Ω)
	fmt.Printf("Привет, %s\n", ответ)
}
