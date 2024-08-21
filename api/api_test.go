package api

import (
	"fmt"
	"testing"
)

func TestCallOpenAI(t *testing.T) {
	tests := []struct {
		query       string
		expected    string
		expectedErr bool
	}{
		{"this is a test string, please reply with 'hello world'", "", false},
	}
	InitOpenAiAPI("", "", "")

	fmt.Println("testing callOpenAI")

	for _, test := range tests {
		result, err := CallOpenAI(test.query)

		if err != nil {
			t.Errorf("err: %v \n", err)
		}

		fmt.Println(result)
	}

}
