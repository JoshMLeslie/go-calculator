package engine

import (
	"testing"
)

func TestApplyOperation(t *testing.T) {
	tests := []struct {
		a, b      float64
		op        string
		expected  float64
		shouldErr bool
	}{
		{3, 2, "+", 5, false},
		{3, 2, "-", 1, false},
		{3, 2, "*", 6, false},
		{3, 2, "/", 1.5, false},
		{3, 0, "/", 0, true},
		{2, 3, "^", 8, false},
		{0, 4, SQUARE_LABEL, 2, false},
		{3, 2, "unknown", 0, true},
	}

	for _, test := range tests {
		result, err := ApplyOperation(test.a, test.b, test.op)
		if (err != nil) != test.shouldErr {
			t.Errorf("ApplyOperation(%f, %f, %s) error = %v, shouldErr %v", test.a, test.b, test.op, err, test.shouldErr)
		}
		if result != test.expected {
			t.Errorf("ApplyOperation(%f, %f, %s) = %f; want %f", test.a, test.b, test.op, result, test.expected)
		}
	}
}

func TestEvaluateRPN(t *testing.T) {
	tests := []struct {
		input     interface{}
		expected  float64
		expectErr bool
	}{
		{"3 4 +", 7, false},
		{"10 5 /", 2, false},
		{"2 3 ^", 8, false},
		{"4 " + SQUARE_LABEL, 2, false},
		{"4 0 /", 0, true},
		{"2 +", 0, true},
	}

	for _, test := range tests {
		result, err := EvaluateRPN(test.input)
		if (err != nil) != test.expectErr {
			t.Errorf("EvaluateRPN(%v) error = %v, expectErr %v", test.input, err, test.expectErr)
		}
		if result != test.expected {
			t.Errorf("EvaluateRPN(%v) = %f; want %f", test.input, result, test.expected)
		}
	}
}

func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		expression string
		expected   []string
		expectErr  bool
	}{
		{"3 + 4", []string{"3", "4", "+"}, false},
		{"10 / 5", []string{"10", "5", "/"}, false},
		{"2 ^ 3", []string{"2", "3", "^"}, false},
		{"( 2 + 3 ) * 4", []string{"2", "3", "+", "4", "*"}, false},
		{"(2+3)*4", []string{"2", "3", "+", "4", "*"}, false},
		{"3 + ( 4 * 2", nil, true},
		{"3 + 4 ) * 2", nil, true},
	}

	for _, test := range tests {
		result, err := InfixToPostfix(test.expression)
		if (err != nil) != test.expectErr {
			t.Errorf("InfixToPostfix(%s) error = %v, expectErr %v", test.expression, err, test.expectErr)
		}
		if !test.expectErr && !equal(result, test.expected) {
			t.Errorf("InfixToPostfix(%s) = %v; want %v", test.expression, result, test.expected)
		}
	}
}

func TestEvalBasicAlgebra(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"3 + 4", 7, false},
		{"10 / 5", 2, false},
		{"2 ^ 3", 8, false},
		{"4 " + SQUARE_LABEL, 2, false},
		{"2 + ( 3 * 4 )", 14, false},
		{"3 + ( 4 * 2", 0, true},
	}

	for _, test := range tests {
		result, err := EvalAlgebra(test.expression)
		if (err != nil) != test.expectErr {
			t.Errorf("EvalBasicAlgebra(%s) error = %v, expectErr %v", test.expression, err, test.expectErr)
		}
		if result != test.expected {
			t.Errorf("EvalBasicAlgebra(%s) = %f; want %f", test.expression, result, test.expected)
		}
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		expression  string
		processType PROCESS_TYPE
		expected    float64
		expectErr   bool
	}{
		{"3 + 4", ALGEBRAIC_BASIC, 7, false},
		{"2 3 +", RPN, 5, false},
		{"2 3 *", RPN, 6, false},
		{"10 5 /", RPN, 2, false},
		{"2 ^ 3", ALGEBRAIC_BASIC, 8, false},
		{"4 " + SQUARE_LABEL, ALGEBRAIC_BASIC, 2, false},
		{"2 + ( 3 * 4 )", ALGEBRAIC_BASIC, 14, false},
	}

	for _, test := range tests {
		result, _ := Calculate(test.expression, test.processType)
		if result != test.expected {
			t.Errorf("Calculate(%s, %d) = %f; want %f", test.expression, test.processType, result, test.expected)
		}
	}
}

// Helper function to compare slices
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
