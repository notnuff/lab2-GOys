package implementation_test

import (
	"errors"
	"fmt"
	"lab2-GOys/implementation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementation(t *testing.T) {
	tests := []struct {
		name        string
		postfix     string
		expectedPre string
		expectedErr error
	}{
		{
			name:        "Valid expression: A B +",
			postfix:     "A B +",
			expectedPre: "+ A B",
			expectedErr: nil,
		},
		{
			name:        "Valid expression: A B C * +",
			postfix:     "A B C * +",
			expectedPre: "+ A * B C",
			expectedErr: nil,
		},
		{
			name:        "Valid expression: A B C * + D +",
			postfix:     "A B C * + D +",
			expectedPre: "+ + A * B C D",
			expectedErr: nil,
		},
		{
			name:        "Valid expression: A B + C D + *",
			postfix:     "A B + C D + *",
			expectedPre: "* + A B + C D",
			expectedErr: nil,
		},
		{
			name:        "Empty expression",
			postfix:     "",
			expectedPre: "",
			expectedErr: errors.New("invalid postfix: empty expression"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := implementation.PostToPre(tt.postfix)

			if tt.expectedErr != nil {
				assert.Error(t, err, "Expected error")
				assert.EqualError(t, err, tt.expectedErr.Error(), "Unexpected error message")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			assert.Equal(t, tt.expectedPre, result, "Unexpected result")
		})
	}
}

func ExamplePostToPre() {
	postfix := "A B + C *"
	result, err := implementation.PostToPre(postfix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Prefix Expression:", result)
	// Output:
	// Prefix Expression: * + A B C
}
