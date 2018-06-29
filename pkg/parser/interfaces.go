package parser

import (
	"bufio"

	"github.com/openshift/junitreport/pkg/api"
)

// TestOutputParser knows how to parse test output to create a collection of test suites
type TestOutputParser interface {
	Parse(input *bufio.Scanner) (*api.TestSuites, error)
}