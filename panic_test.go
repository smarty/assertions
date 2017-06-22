package assertions

import "fmt"

func (fixture *AssertionsFixture) TestShouldPanic() {
	fixture.fail(so(func() {}, ShouldPanic, 1), "This assertion requires exactly 0 comparison values (you provided 1).")
	fixture.fail(so(func() {}, ShouldPanic, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")

	fixture.fail(so(1, ShouldPanic), shouldUseVoidNiladicFunction)
	fixture.fail(so(func(i int) {}, ShouldPanic), shouldUseVoidNiladicFunction)
	fixture.fail(so(func() int { panic("hi") }, ShouldPanic), shouldUseVoidNiladicFunction)

	fixture.fail(so(func() {}, ShouldPanic), shouldHavePanicked)
	fixture.pass(so(func() { panic("hi") }, ShouldPanic))
}

func (fixture *AssertionsFixture) TestShouldNotPanic() {
	fixture.fail(so(func() {}, ShouldNotPanic, 1), "This assertion requires exactly 0 comparison values (you provided 1).")
	fixture.fail(so(func() {}, ShouldNotPanic, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")

	fixture.fail(so(1, ShouldNotPanic), shouldUseVoidNiladicFunction)
	fixture.fail(so(func(i int) {}, ShouldNotPanic), shouldUseVoidNiladicFunction)

	fixture.fail(so(func() { panic("hi") }, ShouldNotPanic), fmt.Sprintf(shouldNotHavePanicked, "hi"))
	fixture.pass(so(func() {}, ShouldNotPanic))
}

func (fixture *AssertionsFixture) TestShouldPanicWith() {
	fixture.fail(so(func() {}, ShouldPanicWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(func() {}, ShouldPanicWith, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(1, ShouldPanicWith, 1), shouldUseVoidNiladicFunction)
	fixture.fail(so(func(i int) {}, ShouldPanicWith, "hi"), shouldUseVoidNiladicFunction)
	fixture.fail(so(func() {}, ShouldPanicWith, "bye"), shouldHavePanicked)
	fixture.fail(so(func() { panic("hi") }, ShouldPanicWith, "bye"), "bye|hi|Expected func() to panic with 'bye' (but it panicked with 'hi')!")

	fixture.pass(so(func() { panic("hi") }, ShouldPanicWith, "hi"))
}

func (fixture *AssertionsFixture) TestShouldNotPanicWith() {
	fixture.fail(so(func() {}, ShouldNotPanicWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(func() {}, ShouldNotPanicWith, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(1, ShouldNotPanicWith, 1), shouldUseVoidNiladicFunction)
	fixture.fail(so(func(i int) {}, ShouldNotPanicWith, "hi"), shouldUseVoidNiladicFunction)
	fixture.fail(so(func() { panic("hi") }, ShouldNotPanicWith, "hi"), "Expected func() NOT to panic with 'hi' (but it did)!")

	fixture.pass(so(func() {}, ShouldNotPanicWith, "bye"))
	fixture.pass(so(func() { panic("hi") }, ShouldNotPanicWith, "bye"))
}
