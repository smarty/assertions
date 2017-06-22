package assertions

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/gunit"
)

/**************************************************************************/

func TestAssertionsFixture(t *testing.T) {
	gunit.Run(new(AssertionsFixture), t)
}

type AssertionsFixture struct {
	*gunit.Fixture
}

func (fixture *AssertionsFixture) Setup() {
	serializer = fixture
}

func (fixture *AssertionsFixture) serialize(expected, actual interface{}, message string) string {
	return fmt.Sprintf("%v|%v|%s", expected, actual, message)
}

func (fixture *AssertionsFixture) serializeDetailed(expected, actual interface{}, message string) string {
	return fmt.Sprintf("%v|%v|%s", expected, actual, message)
}

func (fixture *AssertionsFixture) pass(result string) {
	fixture.Assert(result == success, result)
}

func (fixture *AssertionsFixture) fail(actual string, expected string) {
	actual = format(actual)
	expected = format(expected)

	if actual != expected {
		if actual == "" {
			actual = "(empty)"
		}
		fixture.Errorf("Expected: %s\nActual:   %s\n", expected, actual)
	}
}
func format(message string) string {
	message = strings.Replace(message, "\n", " ", -1)
	for strings.Contains(message, "  ") {
		message = strings.Replace(message, "  ", " ", -1)
	}
	return message
}

/**************************************************************************/

type Thing1 struct {
	a string
}
type Thing2 struct {
	a string
}

type ThingInterface interface {
	Hi()
}

type ThingImplementation struct{}

func (fixture *ThingImplementation) Hi() {}

type IntAlias int
type StringAlias string
type StringSliceAlias []string
type StringStringMapAlias map[string]string

/**************************************************************************/
