package assertions

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func pass(t *testing.T, result string) {
	if result != success {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("Failure:\n%s:%d\nMessage: '%s'", file, line, result)
	}
}

func fail(t *testing.T, actual string, expected string) {
	actual = format(actual)
	expected = format(expected)

	if actual != expected {
		if actual == "" {
			actual = "(empty)"
		}
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("\n%s:%d\nExpected: %s\nActual:   %s\n",
			file, line, expected, actual)
	}
}
func format(message string) string {
	message = strings.Replace(message, "\n", " ", -1)
	for strings.Contains(message, "  ") {
		message = strings.Replace(message, "  ", " ", -1)
	}
	return message
}

type Thing1 struct {
	a string
}
type Thing2 struct {
	a string
}

type Thinger interface {
	Hi()
}

type Thing struct{}

func (self *Thing) Hi() {}

type IntAlias int
type StringAlias string
type StringSliceAlias []string
type StringStringMapAlias map[string]string

/******** FakeSerializer ********/

func init() {
	serializer = newFakeSerializer()
}

type fakeSerializer struct{}

func (self *fakeSerializer) serialize(expected, actual interface{}, message string) string {
	return fmt.Sprintf("%v|%v|%s", expected, actual, message)
}

func (self *fakeSerializer) serializeDetailed(expected, actual interface{}, message string) string {
	return fmt.Sprintf("%v|%v|%s", expected, actual, message)
}

func newFakeSerializer() *fakeSerializer {
	return new(fakeSerializer)
}
