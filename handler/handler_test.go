package handler_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"lab2-GOys/handler"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:     "ValidExpression",
			input:    "5 5 +",
			expected: "+ 5 5",
		},
		{
			name:        "InvalidExpression",
			input:       "5 5 + -",
			expectedErr: errors.New("error converting postfix to prefix: invalid postfix: insufficient operands"),
		},
		{
			name:        "EmptyInput",
			input:       " ",
			expectedErr: errors.New("error converting postfix to prefix: invalid postfix: too many operands"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputReader := strings.NewReader(test.input)
			outputBuffer := new(bytes.Buffer)

			handler := handler.ComputeHandler{
				Input:  inputReader,
				Output: outputBuffer,
			}

			err := handler.Compute()

			if test.expectedErr != nil {
				assert.Error(t, err, "Expected error")
				assert.EqualError(t, err, test.expectedErr.Error(), "Unexpected error message")
				return
			}

			assert.NoError(t, err, "Unexpected error")
			assert.Equal(t, test.expected, outputBuffer.String(), "Unexpected result")
		})
	}
}
