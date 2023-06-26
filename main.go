package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"calculator-project/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter an operation (add, subtract, multiply, divide) or 'q' to quit:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		fmt.Print("Enter the first number: ")
		firstNumStr, _ := reader.ReadString('\n')
		firstNumStr = strings.TrimSpace(firstNumStr)
		firstNum, _ := strconv.Atoi(firstNumStr)

		fmt.Print("Enter the second number: ")
		secondNumStr, _ := reader.ReadString('\n')
		secondNumStr = strings.TrimSpace(secondNumStr)
		secondNum, _ := strconv.Atoi(secondNumStr)

		switch input {
		case "add":
			result := calculator.Add(firstNum, secondNum)
			fmt.Printf("Result: %d\n", result)
		case "subtract":
			result := calculator.Subtract(firstNum, secondNum)
			fmt.Printf("Result: %d\n", result)
		case "multiply":
			result := calculator.Multiply(firstNum, secondNum)
			fmt.Printf("Result: %d\n", result)
		case "divide":
			result := calculator.Divide(firstNum, secondNum)
			fmt.Printf("Result: %d\n", result)
		default:
			fmt.Println("Invalid operation")
		}
	}
}
