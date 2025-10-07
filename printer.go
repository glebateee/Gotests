package main

import "fmt"

// Print просто выводит сообщение
func Print(msg string) string {
	return fmt.Sprintf("Message: %s", msg)
}

// PrintWithPrefix выводит сообщение с префиксом
func PrintWithPrefix(prefix, msg string) string {
	return fmt.Sprintf("[%s] %s", prefix, msg)
}

// PrintNumbers форматирует числа
func PrintNumbers(nums ...int) string {
	return fmt.Sprintf("Numbers: %v", nums)
}
