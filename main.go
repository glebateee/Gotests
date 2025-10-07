package main

import "fmt"

func main() {
	fmt.Println(Print("Hello from Gotests!"))
	fmt.Println(PrintWithPrefix("DEBUG", "Application started"))
	fmt.Println(PrintNumbers(10, 20, 30))
}
