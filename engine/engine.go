package engine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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
			if len(stack) < 1 && (token == SQUARE) {
				return 0, errors.New("invalid expression: not enough operands for square root")
			} else if len(stack) < 2 && (token != SQUARE) {
				return 0, errors.New("invalid expression: not enough operands")
			}
			var result float64
			var err error
			if token == SQUARE {
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result, err = ApplyOperation(0, b, token)
				if err != nil {
					return 0, err
				}
			} else {
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]

				result, err = ApplyOperation(a, b, token)
				if err != nil {
					return 0, err
				}
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
func InfixToPostfix(exp string) ([]string, error) {
	var err error
	parenOpen := false
	opStack := StringStack{}
	result := []string{}

tokenLoop:
	for _, char := range tokenize(exp) {
		_, numErr := strconv.Atoi(char)

		if !isOperator(char) && numErr == nil {
			result = append(result, char)
		} else if isOpeningParenthesis(char) {
			parenOpen = true
			opStack.Push(char)
		} else if isClosingParenthesis(char) {
			if parenOpen {
				parenOpen = false
			} else {
				err = errors.New("unopened parenthesis")
			}
			for !opStack.IsEmpty() && !isMatchingParenthesis(opStack.Top(), char) {
				result = append(result, opStack.Pop())
			}
			opStack.Pop() // pop the matching close paren
		} else if isOperator(char) {
			top := opStack.Top()
			higherPrecedence := hasHigherPrecedence(top, char)
			openingParen := isOpeningParenthesis(top)

			for !opStack.IsEmpty() &&
				higherPrecedence &&
				!openingParen {
				result = append(result, opStack.Pop())
			}
			opStack.Push(char)
		} else if err != nil {
			err = errors.New("err testing num: " + char)
		} else {
			err = errors.New("unknown token in expresion: " + char)
		}
		if err != nil {
			break tokenLoop
		}
	}

	if parenOpen {
		err = errors.New("unclosed parenthesis")
	}

	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	// dump remaining operators onto result, if any
	if !opStack.IsEmpty() {
		result = append(result, opStack.items...)
	}

	return result, nil
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
