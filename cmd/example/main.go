package main

import (
	"flag"
	"fmt"
	"io"
	"lab2-GOys/handler"
	"log"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File to read")
	outputFile      = flag.String("o", "", "File to output")
)

type StringReaderCloser struct {
	*strings.Reader
}

func (s *StringReaderCloser) Close() error {
	return nil
}

func (s *StringReaderCloser) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("Write operation not supported")
}

func NewStringReaderCloser(s string) *StringReaderCloser {
	return &StringReaderCloser{strings.NewReader(s)}
}

func main() {
	flag.Parse()
	flag.CommandLine.SetOutput(os.Stderr)

	if len(*inputExpression) < 1 && len(*inputFile) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var (
		readWriterCloser io.ReadWriteCloser
		err              error
	)

	switch {
	case len(*inputExpression) > 0:
		readWriterCloser = NewStringReaderCloser(*inputExpression)
	case len(*inputFile) > 0:
		readWriterCloser, err = os.Open(*inputFile)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer func() {
			if cerr := readWriterCloser.Close(); cerr != nil {
				log.Fatalf("Error closing file: %v", cerr)
			}
		}()
	}

	var outputWriter io.Writer = os.Stdout
	if len(*outputFile) > 0 {
		outputFile, err := os.Create(*outputFile)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer func() {
			if cerr := outputFile.Close(); cerr != nil {
				log.Fatalf("Error closing output file: %v", cerr)
			}
		}()
		outputWriter = outputFile
	}

	handler := handler.ComputeHandler{
		Input:  readWriterCloser,
		Output: outputWriter,
	}

	if err := handler.Compute(); err != nil {
		log.Fatal(err)
	}
}
