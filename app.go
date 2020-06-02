package main

import "fmt"

func main() {
	var op string
	var num1, num2 int
	fmt.Print("Please write what do you want to do: ")
	fmt.Scanln(&num1, &op, &num2)
	output := 0
	if op == "+" {
		output = num1 + num2
	}
	if op == "_" {
		output = num1 - num2
	}
	if op == "*" {
		output = num1 * num2
	}
	if op == "/" {
		output = num1 / num2
	}
	if op == "◦/◦" {
		output = num1 % num2
	}

	fmt.Printf("%d %s %d = %d\t\n", num1, num1, num2, output)
}
