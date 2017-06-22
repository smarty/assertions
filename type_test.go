package assertions

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func (fixture *AssertionsFixture) TestShouldHaveSameTypeAs() {
	fixture.fail(so(1, ShouldHaveSameTypeAs), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(1, ShouldHaveSameTypeAs, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(nil, ShouldHaveSameTypeAs, 0), "int|<nil>|Expected '<nil>' to be: 'int' (but was: '<nil>')!")
	fixture.fail(so(1, ShouldHaveSameTypeAs, "asdf"), "string|int|Expected '1' to be: 'string' (but was: 'int')!")

	fixture.pass(so(1, ShouldHaveSameTypeAs, 0))
	fixture.pass(so(nil, ShouldHaveSameTypeAs, nil))
}

func (fixture *AssertionsFixture) TestShouldNotHaveSameTypeAs() {
	fixture.fail(so(1, ShouldNotHaveSameTypeAs), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(1, ShouldNotHaveSameTypeAs, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(1, ShouldNotHaveSameTypeAs, 0), "Expected '1' to NOT be: 'int' (but it was)!")
	fixture.fail(so(nil, ShouldNotHaveSameTypeAs, nil), "Expected '<nil>' to NOT be: '<nil>' (but it was)!")

	fixture.pass(so(nil, ShouldNotHaveSameTypeAs, 0))
	fixture.pass(so(1, ShouldNotHaveSameTypeAs, "asdf"))
}

func (fixture *AssertionsFixture) TestShouldImplement() {
	var ioReader *io.Reader
	var response http.Response = http.Response{}
	var responsePtr *http.Response = new(http.Response)
	var reader = bytes.NewBufferString("")

	fixture.fail(so(reader, ShouldImplement), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(reader, ShouldImplement, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 2).")
	fixture.fail(so(reader, ShouldImplement, ioReader, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(reader, ShouldImplement, "foo"), shouldCompareWithInterfacePointer)
	fixture.fail(so(reader, ShouldImplement, 1), shouldCompareWithInterfacePointer)
	fixture.fail(so(reader, ShouldImplement, nil), shouldCompareWithInterfacePointer)

	fixture.fail(so(nil, ShouldImplement, ioReader), shouldNotBeNilActual)
	fixture.fail(so(1, ShouldImplement, ioReader), "Expected: 'io.Reader interface support'\nActual:   '*int' does not implement the interface!")

	fixture.fail(so(response, ShouldImplement, ioReader), "Expected: 'io.Reader interface support'\nActual:   '*http.Response' does not implement the interface!")
	fixture.fail(so(responsePtr, ShouldImplement, ioReader), "Expected: 'io.Reader interface support'\nActual:   '*http.Response' does not implement the interface!")
	fixture.pass(so(reader, ShouldImplement, ioReader))
	fixture.pass(so(reader, ShouldImplement, (*io.Reader)(nil)))
}

func (fixture *AssertionsFixture) TestShouldNotImplement() {
	var ioReader *io.Reader
	var response http.Response = http.Response{}
	var responsePtr *http.Response = new(http.Response)
	var reader io.Reader = bytes.NewBufferString("")

	fixture.fail(so(reader, ShouldNotImplement), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(reader, ShouldNotImplement, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 2).")
	fixture.fail(so(reader, ShouldNotImplement, ioReader, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(reader, ShouldNotImplement, "foo"), shouldCompareWithInterfacePointer)
	fixture.fail(so(reader, ShouldNotImplement, 1), shouldCompareWithInterfacePointer)
	fixture.fail(so(reader, ShouldNotImplement, nil), shouldCompareWithInterfacePointer)

	fixture.fail(so(reader, ShouldNotImplement, ioReader), "Expected         '*bytes.Buffer'\nto NOT implement   'io.Reader' (but it did)!")
	fixture.fail(so(nil, ShouldNotImplement, ioReader), shouldNotBeNilActual)
	fixture.pass(so(1, ShouldNotImplement, ioReader))
	fixture.pass(so(response, ShouldNotImplement, ioReader))
	fixture.pass(so(responsePtr, ShouldNotImplement, ioReader))
}

func (fixture *AssertionsFixture) TestShouldBeError() {
	fixture.fail(so(nil, ShouldBeError, "too", "many"), "This assertion allows 1 or fewer comparison values (you provided 2).")

	fixture.fail(so(1, ShouldBeError), "Expected an error value (but was 'int' instead)!")
	fixture.fail(so(nil, ShouldBeError), "Expected an error value (but was '<nil>' instead)!")

	error1 := errors.New("Message")

	fixture.fail(so(error1, ShouldBeError, 42), "The final argument to this assertion must be a string or an error value (you provided: 'int').")
	fixture.fail(so(error1, ShouldBeError, "Wrong error message"), "Wrong error message|Message|Expected: 'Wrong error message' Actual: 'Message' (Should be equal)")

	fixture.pass(so(error1, ShouldBeError))
	fixture.pass(so(error1, ShouldBeError, error1))
	fixture.pass(so(error1, ShouldBeError, error1.Error()))
}
