package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type PROCESS_TYPE int

const (
	ALGEBRAIC_BASIC PROCESS_TYPE = iota
	ALGEBRAIC_ADVANCED
	RPN
	MATHML
)

// Precedence of operators
var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

// ApplyOperation performs the arithmetic operation between two operands
func ApplyOperation(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	case "^":
		return float64(int(a) ^ int(b)), nil
	default:
		return 0, errors.New("unknown operator")
	}
}

// EvaluateRPN evaluates a reverse Polish notation expression and returns the result.
func EvaluateRPN(input interface{}) (float64, error) {
	var tokens []string
	stack := []float64{}
	str, ok := input.(string)
	if ok {
		tokens = strings.Split(str, " ")
	} else if tokens, ok = input.([]string); !ok {
		// NB sets tokens to []string else PrintsLn !ok
		fmt.Println("Unsupported type")
	}

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression: not enough operands")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result, err := ApplyOperation(a, b, token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression: too many operands")
	}
	return stack[0], nil
}

// InfixToPostfix converts an infix expression to postfix (RPN)
func InfixToPostfix(expression string) ([]string, error) {
	var output []string
	var operators []string

	tokens := strings.FieldsFunc(expression, func(r rune) bool {
		return unicode.IsSpace(r) || strings.ContainsRune("+-*/^()", r)
	})

	for _, token := range tokens {
		switch {
		case unicode.IsDigit(rune(token[0])):
			output = append(output, token)
		case token == "(":
			operators = append(operators, token)
		case token == ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1] // pop '('
		default:
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

// EvaluateExpression evaluates a given algebraic expression string
func EvalBasicAlgebra(expression string) (float64, error) {
	postfix, err := InfixToPostfix(expression)
	if err != nil {
		return 0, err
	}
	return EvaluateRPN(postfix)
}

func Calculate(
	expression string,
	processType PROCESS_TYPE,
) float64 {
	var (
		result float64
		err    error
	)
	switch processType {
	case ALGEBRAIC_BASIC:
		result, err = EvalBasicAlgebra(expression)
	case ALGEBRAIC_ADVANCED:
		return 0
	case RPN:
		result, err = EvaluateRPN(expression)
	case MATHML:
		return 0
	}
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	return result
}
