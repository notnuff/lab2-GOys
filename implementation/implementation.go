package implementation

import (
	"errors"
	"fmt"
	"strings"
)

func IsOperator(char string) bool {
	return char == "+" || char == "-" || char == "*" || char == "/"
}

func PostToPre(postfix string) (string, error) {
	if len(postfix) == 0 {
		return "", errors.New("invalid postfix: empty expression")
	}

	stack := []string{}

	tokens := strings.Fields(postfix)
	for _, token := range tokens {
		if IsOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("invalid postfix: insufficient operands")
			}

			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			newExpr := fmt.Sprintf("%s %s %s", token, operand1, operand2)
			stack = append(stack, newExpr)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid postfix: too many operands")
	}

	return stack[0], nil
}
