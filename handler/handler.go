package handler

import (
	"fmt"
	"io"
	"lab2-GOys/implementation"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Our handler
func (ch *ComputeHandler) Compute() error {
	postfixExp, err := readFromReader(ch.Input)
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	prefixExp, err := implementation.PostToPre(postfixExp)
	if err != nil {
		return fmt.Errorf("error converting postfix to prefix: %v", err)
	}

	_, err = ch.Output.Write([]byte(prefixExp))
	if err != nil {
		return fmt.Errorf("error writing to output: %v", err)
	}

	return nil
}

func readFromReader(reader io.Reader) (string, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
