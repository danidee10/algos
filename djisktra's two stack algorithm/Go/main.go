/* Implementation of djisktra's two stack algorithm */

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type stack []string

var operandStack stack
var operatorStack stack

func (s *stack) Push(value string) {
	*s = append(*s, value)
}

func (s *stack) Pop() (string, error) {
	st := *s

	if len(st) == 0 {
		return "", errors.New("Pop from empty stack")
	}

	lastPosition := len(st) - 1
	value := st[lastPosition]

	// delete the value in the stack
	*s = append(st[:lastPosition], st[lastPosition+1:]...)

	return value, nil
}

func applyOperator(val1 int, val2 int, operator string) {

	if operator == "*" {
		result := val1 * val2

		operandStack.Push(strconv.Itoa(result))
	}

	if operator == "/" {
		result := val2 / val1

		operandStack.Push(strconv.Itoa(result))
	}

	if operator == "+" {
		result := val1 + val2

		operandStack.Push(strconv.Itoa(result))
	}

	if operator == "-" {
		result := val1 - val2

		operandStack.Push(strconv.Itoa(result))
	}
}

func main() {

	inputString := "( ( 1 + 2 ) * ( 3 * 3 ) ) + 2"

	// Remove spaces
	inputString = strings.Replace(inputString, " ", "", -1)

	for _, rune := range inputString {
		char := string(rune)

		if char == "+" || char == "*" || char == "-" || char == "/" {
			operatorStack.Push(char)
		} else {
			if char != "(" && char != ")" {
				operandStack.Push(char)
			}
		}

		if char == ")" {
			operand1, _ := operandStack.Pop()
			operand2, _ := operandStack.Pop()

			val2, _ := strconv.Atoi(operand1)
			val1, _ := strconv.Atoi(operand2)
			operator, _ := operatorStack.Pop()

			applyOperator(val1, val2, operator)
		}
	}

	for {
		operator, err := operatorStack.Pop()
		if err != nil {
			fmt.Println(operandStack[0])
			break
		}

		operand1, _ := operandStack.Pop()
		operand2, _ := operandStack.Pop()

		val2, _ := strconv.Atoi(operand1)
		val1, _ := strconv.Atoi(operand2)

		applyOperator(val1, val2, operator)
	}
}
