package engine

import (
	"errors"
	"math"
	"unicode"
)

type PROCESS_TYPE int

const (
	ALGEBRAIC_BASIC PROCESS_TYPE = iota
	ALGEBRAIC_ADVANCED
	RPN
	MATHML
)

const SQUARE string = "sqr"

var operators = map[string]bool{
	"+":    true,
	"-":    true,
	"*":    true,
	"/":    true,
	"^":    true,
	"(":    true,
	")":    true,
	SQUARE: true,
}

func isOperator(char string) bool {
	return operators[char]
}

// StringStack
type StringStack struct {
	items []string
}

func (s *StringStack) Push(item string) {
	s.items = append(s.items, item)
}
func (s *StringStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Top returns the element at the top of the stack without removing it
func (s *StringStack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.items[len(s.items)-1]
}

// Pop removes and returns the element at the top of the stack
func (s *StringStack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	// Get the top element
	top := s.items[len(s.items)-1]
	// Remove the top element from the stack
	s.items = s.items[:len(s.items)-1]
	return top
}

// end StringStack
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
		return math.Pow(a, b), nil
	case SQUARE:
		return math.Sqrt(b), nil
	default:
		return 0, errors.New("unknown operator")
	}
}

func isOpeningParenthesis(char string) bool {
	switch char {
	case "(":
		return true
	case "{":
		return true
	case "[":
		return true
	default:
		return false
	}
}

func isClosingParenthesis(char string) bool {
	switch char {
	case ")":
		return true
	case "}":
		return true
	case "]":
		return true
	default:
		return false
	}
}

func isMatchingParenthesis(opening, closing string) bool {
	switch opening {
	case "(":
		return closing == ")"
	case "{":
		return closing == "}"
	case "[":
		return closing == "]"
	default:
		return false
	}
}

func hasHigherPrecedence(target, source string) bool {
	return (target == "*" || target == "/") && (source == "+" || source == "-")
}

// Tokenize the equation by handling numbers, operators, and parentheses
func tokenize(exp string) []string {
	var tokens []string
	var currentToken string

	for _, char := range exp {
		switch {
		case unicode.IsDigit(char) || char == '.':
			// If the character is a digit or a decimal point, add it to the current token
			currentToken += string(char)
		case char == '(' || char == ')' || char == '+' || char == '-' || char == '*' || char == '/' || char == '^':
			// If the current token is not empty, add it to tokens
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			// Add the operator or parenthesis as a separate token
			tokens = append(tokens, string(char))
		default:
			// If the character doesn't match any expected type, reset the current token
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
		}
	}

	// Add the last token if there's any
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}
