package assert

import (
	"fmt"
	"log"
	"runtime"
)

// Result contains a single assertion failure as an error.
// You should not create a Result directly, use So instead.
// Once created, a Result is read-only and only allows
// queries using the provided methods.
type Result struct {
	err error
}

// So is a convenience function (as opposed to an inconvenience function?)
// for running assertions on arbitrary arguments in any context. It allows you to perform
// assertion-like behavior and decide what happens in the event of a failure.
// It is a variant of assertions.So in ever respect except its return value.
//
// Examples:
//
//    assert.So(1, should.Equal, 1).Println() // Calls fmt.Print with the failure message and file:line header.
//    assert.So(1, should.Equal, 1).Log()     // Calls log.Print with the failure message and file:line header.
//    assert.So(1, should.Equal, 1).Panic()   // Calls log.Panic with the failure message and file:line header.
//    assert.So(1, should.Equal, 1).Fatal()   // Calls log.Fatal with the failure message and file:line header.
//    if err := assert.So(1, should.Equal, 1).Error(); err != nil {
//        // Handle the error, which will include the failure message and file:line header.
//    }
//
func So(actual interface{}, assert func(interface{}, ...interface{}) string, expected ...interface{}) *Result {
	result := new(Result)
	failure := assert(actual, expected...)
	if len(failure) > 0 {
		_, file, line, _ := runtime.Caller(1)
		result.err = fmt.Errorf("Assertion failure at %s:%d\n%s", file, line, failure)
	}
	return result
}

// Failed returns true if the assertion failed, false if it passed.
func (this *Result) Failed() bool {
	return !this.Passed()
}

// Passed returns true if the assertion passed, false if it failed.
func (this *Result) Passed() bool {
	return this.err == nil
}

// Returns the error representing an assertion failure, which is nil in the case of a passed assertion.
func (this *Result) Error() error {
	return this.err
}

// String implements fmt.Stringer. It returns the error as a string in the case of an assertion failure.
func (this *Result) String() string {
	if this.Failed() {
		return this.err.Error()
	}
	return ""
}

// Println calls fmt.Println in the case of an assertion failure.
func (this *Result) Println() {
	if this.Failed() {
		fmt.Println(this)
	}
}

// Log calls log.Print in the case of an assertion failure.
func (this *Result) Log() {
	if this.Failed() {
		log.Print(this)
	}
}

// Panic calls log.Panic in the case of an assertion failure.
func (this *Result) Panic() {
	if this.Failed() {
		log.Panic(this)
	}
}

// Fatal calls log.Fatal in the case of an assertion failure.
func (this *Result) Fatal() {
	if this.Failed() {
		log.Fatal(this)
	}
}

// assertion is a copy of github.com/smartystreets/assertions.assertion.
type assertion func(actual interface{}, expected ...interface{}) string
