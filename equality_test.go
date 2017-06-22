package assertions

import (
	"fmt"
	"reflect"
)

func (fixture *AssertionsFixture) TestShouldEqual() {
	fixture.fail(so(1, ShouldEqual), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(1, ShouldEqual, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	fixture.fail(so(1, ShouldEqual, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.pass(so(1, ShouldEqual, 1))
	fixture.fail(so(1, ShouldEqual, 2), "2|1|Expected: '2' Actual: '1' (Should be equal)")
	fixture.fail(so(1, ShouldEqual, "1"), "1|1|Expected: '1' (string) Actual: '1' (int) (Should be equal, type mismatch)")

	fixture.pass(so(true, ShouldEqual, true))
	fixture.fail(so(true, ShouldEqual, false), "false|true|Expected: 'false' Actual: 'true' (Should be equal)")

	fixture.pass(so("hi", ShouldEqual, "hi"))
	fixture.fail(so("hi", ShouldEqual, "bye"), "bye|hi|Expected: 'bye' Actual: 'hi' (Should be equal)")

	fixture.pass(so(42, ShouldEqual, uint(42)))

	fixture.fail(so(Thing1{"hi"}, ShouldEqual, Thing1{}), "{}|{hi}|Expected: '{}' Actual: '{hi}' (Should be equal)")
	fixture.fail(so(Thing1{"hi"}, ShouldEqual, Thing1{"hi"}), "{hi}|{hi}|Expected: '{hi}' Actual: '{hi}' (Should be equal)")
	fixture.fail(so(&Thing1{"hi"}, ShouldEqual, &Thing1{"hi"}), "&{hi}|&{hi}|Expected: '&{hi}' Actual: '&{hi}' (Should be equal)")

	fixture.fail(so(Thing1{}, ShouldEqual, Thing2{}), "{}|{}|Expected: '{}' Actual: '{}' (Should be equal)")
}

func (fixture *AssertionsFixture) TestShouldNotEqual() {
	fixture.fail(so(1, ShouldNotEqual), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(1, ShouldNotEqual, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	fixture.fail(so(1, ShouldNotEqual, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.pass(so(1, ShouldNotEqual, 2))
	fixture.pass(so(1, ShouldNotEqual, "1"))
	fixture.fail(so(1, ShouldNotEqual, 1), "Expected '1' to NOT equal '1' (but it did)!")

	fixture.pass(so(true, ShouldNotEqual, false))
	fixture.fail(so(true, ShouldNotEqual, true), "Expected 'true' to NOT equal 'true' (but it did)!")

	fixture.pass(so("hi", ShouldNotEqual, "bye"))
	fixture.fail(so("hi", ShouldNotEqual, "hi"), "Expected 'hi' to NOT equal 'hi' (but it did)!")

	fixture.pass(so(&Thing1{"hi"}, ShouldNotEqual, &Thing1{"hi"}))
	fixture.pass(so(Thing1{"hi"}, ShouldNotEqual, Thing1{"hi"}))
	fixture.pass(so(Thing1{}, ShouldNotEqual, Thing1{}))
	fixture.pass(so(Thing1{}, ShouldNotEqual, Thing2{}))
}

func (fixture *AssertionsFixture) TestShouldAlmostEqual() {
	fixture.fail(so(1, ShouldAlmostEqual), "This assertion requires exactly one comparison value and an optional delta (you provided neither)")
	fixture.fail(so(1, ShouldAlmostEqual, 1, 2, 3), "This assertion requires exactly one comparison value and an optional delta (you provided more values)")
	fixture.fail(so(1, ShouldAlmostEqual, "1"), "The comparison value must be a numerical type, but was: string")
	fixture.fail(so(1, ShouldAlmostEqual, 1, "1"), "The delta value must be a numerical type, but was: string")
	fixture.fail(so("1", ShouldAlmostEqual, 1), "The actual value must be a numerical type, but was: string")

	// with the default delta
	fixture.pass(so(.99999999999999, ShouldAlmostEqual, uint(1)))
	fixture.pass(so(1, ShouldAlmostEqual, .99999999999999))
	fixture.pass(so(1.3612499999999996, ShouldAlmostEqual, 1.36125))
	fixture.pass(so(0.7285312499999999, ShouldAlmostEqual, 0.72853125))
	fixture.fail(so(1, ShouldAlmostEqual, .99), "Expected '1' to almost equal '0.99' (but it didn't)!")

	// with a different delta
	fixture.pass(so(100.0, ShouldAlmostEqual, 110.0, 10.0))
	fixture.fail(so(100.0, ShouldAlmostEqual, 111.0, 10.5), "Expected '100' to almost equal '111' (but it didn't)!")

	// various ints should work
	fixture.pass(so(100, ShouldAlmostEqual, 100.0))
	fixture.pass(so(int(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(int8(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(int16(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(int32(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(int64(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(uint(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(uint8(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(uint16(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(uint32(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(uint64(100), ShouldAlmostEqual, 100.0))
	fixture.pass(so(100, ShouldAlmostEqual, 100.0))
	fixture.fail(so(100, ShouldAlmostEqual, 99.0), "Expected '100' to almost equal '99' (but it didn't)!")

	// floats should work
	fixture.pass(so(float64(100.0), ShouldAlmostEqual, float32(100.0)))
	fixture.fail(so(float32(100.0), ShouldAlmostEqual, 99.0, float32(0.1)), "Expected '100' to almost equal '99' (but it didn't)!")
}

func (fixture *AssertionsFixture) TestShouldNotAlmostEqual() {
	fixture.fail(so(1, ShouldNotAlmostEqual), "This assertion requires exactly one comparison value and an optional delta (you provided neither)")
	fixture.fail(so(1, ShouldNotAlmostEqual, 1, 2, 3), "This assertion requires exactly one comparison value and an optional delta (you provided more values)")

	// with the default delta
	fixture.fail(so(1, ShouldNotAlmostEqual, .99999999999999), "Expected '1' to NOT almost equal '0.99999999999999' (but it did)!")
	fixture.fail(so(1.3612499999999996, ShouldNotAlmostEqual, 1.36125), "Expected '1.3612499999999996' to NOT almost equal '1.36125' (but it did)!")
	fixture.pass(so(1, ShouldNotAlmostEqual, .99))

	// with a different delta
	fixture.fail(so(100.0, ShouldNotAlmostEqual, 110.0, 10.0), "Expected '100' to NOT almost equal '110' (but it did)!")
	fixture.pass(so(100.0, ShouldNotAlmostEqual, 111.0, 10.5))

	// ints should work
	fixture.fail(so(100, ShouldNotAlmostEqual, 100.0), "Expected '100' to NOT almost equal '100' (but it did)!")
	fixture.pass(so(100, ShouldNotAlmostEqual, 99.0))

	// float32 should work
	fixture.fail(so(float64(100.0), ShouldNotAlmostEqual, float32(100.0)), "Expected '100' to NOT almost equal '100' (but it did)!")
	fixture.pass(so(float32(100.0), ShouldNotAlmostEqual, 99.0, float32(0.1)))
}

func (fixture *AssertionsFixture) TestShouldResemble() {
	fixture.fail(so(Thing1{"hi"}, ShouldResemble), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(Thing1{"hi"}, ShouldResemble, Thing1{"hi"}, Thing1{"hi"}), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.pass(so(Thing1{"hi"}, ShouldResemble, Thing1{"hi"}))
	fixture.fail(so(Thing1{"hi"}, ShouldResemble, Thing1{"bye"}), `{bye}|{hi}|Expected: 'assertions.Thing1{a:"bye"}' Actual: 'assertions.Thing1{a:"hi"}' (Should resemble)!`)

	var (
		a []int
		b []int = []int{}
	)

	fixture.fail(so(a, ShouldResemble, b), `[]|[]|Expected: '[]int{}' Actual: '[]int(nil)' (Should resemble)!`)
	fixture.fail(so(2, ShouldResemble, 1), `1|2|Expected: '1' Actual: '2' (Should resemble)!`)

	fixture.fail(so(StringStringMapAlias{"hi": "bye"}, ShouldResemble, map[string]string{"hi": "bye"}),
		`map[hi:bye]|map[hi:bye]|Expected: 'map[string]string{"hi":"bye"}' Actual: 'assertions.StringStringMapAlias{"hi":"bye"}' (Should resemble)!`)
	fixture.fail(so(StringSliceAlias{"hi", "bye"}, ShouldResemble, []string{"hi", "bye"}),
		`[hi bye]|[hi bye]|Expected: '[]string{"hi", "bye"}' Actual: 'assertions.StringSliceAlias{"hi", "bye"}' (Should resemble)!`)

	// some types come out looking the same when represented with "%#v" so we show type mismatch info:
	fixture.fail(so(StringAlias("hi"), ShouldResemble, "hi"), `hi|hi|Expected: '"hi"' Actual: 'assertions.StringAlias("hi")' (Should resemble)!`)
	fixture.fail(so(IntAlias(42), ShouldResemble, 42), `42|42|Expected: '42' Actual: 'assertions.IntAlias(42)' (Should resemble)!`)
}

func (fixture *AssertionsFixture) TestShouldNotResemble() {
	fixture.fail(so(Thing1{"hi"}, ShouldNotResemble), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(Thing1{"hi"}, ShouldNotResemble, Thing1{"hi"}, Thing1{"hi"}), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.pass(so(Thing1{"hi"}, ShouldNotResemble, Thing1{"bye"}))
	fixture.fail(so(Thing1{"hi"}, ShouldNotResemble, Thing1{"hi"}),
		`Expected '"assertions.Thing1{a:\"hi\"}"' to NOT resemble '"assertions.Thing1{a:\"hi\"}"' (but it did)!`)

	fixture.pass(so(map[string]string{"hi": "bye"}, ShouldResemble, map[string]string{"hi": "bye"}))
	fixture.pass(so(IntAlias(42), ShouldNotResemble, 42))

	fixture.pass(so(StringSliceAlias{"hi", "bye"}, ShouldNotResemble, []string{"hi", "bye"}))
}

func (fixture *AssertionsFixture) TestShouldPointTo() {
	t1 := &Thing1{}
	t2 := t1
	t3 := &Thing1{}

	pointer1 := reflect.ValueOf(t1).Pointer()
	pointer3 := reflect.ValueOf(t3).Pointer()

	fixture.fail(so(t1, ShouldPointTo), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(t1, ShouldPointTo, t2, t3), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.pass(so(t1, ShouldPointTo, t2))
	fixture.fail(so(t1, ShouldPointTo, t3), fmt.Sprintf(
		"%v|%v|Expected '&{a:}' (address: '%v') and '&{a:}' (address: '%v') to be the same address (but their weren't)!",
		pointer3, pointer1, pointer1, pointer3))

	t4 := Thing1{}
	t5 := t4

	fixture.fail(so(t4, ShouldPointTo, t5), "Both arguments should be pointers (the first was not)!")
	fixture.fail(so(&t4, ShouldPointTo, t5), "Both arguments should be pointers (the second was not)!")
	fixture.fail(so(nil, ShouldPointTo, nil), "Both arguments should be pointers (the first was nil)!")
	fixture.fail(so(&t4, ShouldPointTo, nil), "Both arguments should be pointers (the second was nil)!")
}

func (fixture *AssertionsFixture) TestShouldNotPointTo() {
	t1 := &Thing1{}
	t2 := t1
	t3 := &Thing1{}

	pointer1 := reflect.ValueOf(t1).Pointer()

	fixture.fail(so(t1, ShouldNotPointTo), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(t1, ShouldNotPointTo, t2, t3), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.pass(so(t1, ShouldNotPointTo, t3))
	fixture.fail(so(t1, ShouldNotPointTo, t2), fmt.Sprintf("Expected '&{a:}' and '&{a:}' to be different references (but they matched: '%v')!", pointer1))

	t4 := Thing1{}
	t5 := t4

	fixture.fail(so(t4, ShouldNotPointTo, t5), "Both arguments should be pointers (the first was not)!")
	fixture.fail(so(&t4, ShouldNotPointTo, t5), "Both arguments should be pointers (the second was not)!")
	fixture.fail(so(nil, ShouldNotPointTo, nil), "Both arguments should be pointers (the first was nil)!")
	fixture.fail(so(&t4, ShouldNotPointTo, nil), "Both arguments should be pointers (the second was nil)!")
}

func (fixture *AssertionsFixture) TestShouldBeNil() {
	fixture.fail(so(nil, ShouldBeNil, nil, nil, nil), "This assertion requires exactly 0 comparison values (you provided 3).")
	fixture.fail(so(nil, ShouldBeNil, nil), "This assertion requires exactly 0 comparison values (you provided 1).")

	fixture.pass(so(nil, ShouldBeNil))
	fixture.fail(so(1, ShouldBeNil), "Expected: nil Actual: '1'")

	var thing ThingInterface
	fixture.pass(so(thing, ShouldBeNil))
	thing = &ThingImplementation{}
	fixture.fail(so(thing, ShouldBeNil), "Expected: nil Actual: '&{}'")

	var thingOne *Thing1
	fixture.pass(so(thingOne, ShouldBeNil))

	var nilSlice []int
	fixture.pass(so(nilSlice, ShouldBeNil))

	var nilMap map[string]string
	fixture.pass(so(nilMap, ShouldBeNil))

	var nilChannel chan int
	fixture.pass(so(nilChannel, ShouldBeNil))

	var nilFunc func()
	fixture.pass(so(nilFunc, ShouldBeNil))

	var nilInterface interface{}
	fixture.pass(so(nilInterface, ShouldBeNil))
}

func (fixture *AssertionsFixture) TestShouldNotBeNil() {
	fixture.fail(so(nil, ShouldNotBeNil, nil, nil, nil), "This assertion requires exactly 0 comparison values (you provided 3).")
	fixture.fail(so(nil, ShouldNotBeNil, nil), "This assertion requires exactly 0 comparison values (you provided 1).")

	fixture.fail(so(nil, ShouldNotBeNil), "Expected '<nil>' to NOT be nil (but it was)!")
	fixture.pass(so(1, ShouldNotBeNil))

	var thing ThingInterface
	fixture.fail(so(thing, ShouldNotBeNil), "Expected '<nil>' to NOT be nil (but it was)!")
	thing = &ThingImplementation{}
	fixture.pass(so(thing, ShouldNotBeNil))
}

func (fixture *AssertionsFixture) TestShouldBeTrue() {
	fixture.fail(so(true, ShouldBeTrue, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	fixture.fail(so(true, ShouldBeTrue, 1), "This assertion requires exactly 0 comparison values (you provided 1).")

	fixture.fail(so(false, ShouldBeTrue), "Expected: true Actual: false")
	fixture.fail(so(1, ShouldBeTrue), "Expected: true Actual: 1")
	fixture.pass(so(true, ShouldBeTrue))
}

func (fixture *AssertionsFixture) TestShouldBeFalse() {
	fixture.fail(so(false, ShouldBeFalse, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	fixture.fail(so(false, ShouldBeFalse, 1), "This assertion requires exactly 0 comparison values (you provided 1).")

	fixture.fail(so(true, ShouldBeFalse), "Expected: false Actual: true")
	fixture.fail(so(1, ShouldBeFalse), "Expected: false Actual: 1")
	fixture.pass(so(false, ShouldBeFalse))
}

func (fixture *AssertionsFixture) TestShouldBeZeroValue() {
	fixture.fail(so(0, ShouldBeZeroValue, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	fixture.fail(so(false, ShouldBeZeroValue, true), "This assertion requires exactly 0 comparison values (you provided 1).")

	fixture.fail(so(1, ShouldBeZeroValue), "0|1|'1' should have been the zero value")                                       //"Expected: (zero value) Actual: 1")
	fixture.fail(so(true, ShouldBeZeroValue), "false|true|'true' should have been the zero value")                          //"Expected: (zero value) Actual: true")
	fixture.fail(so("123", ShouldBeZeroValue), "|123|'123' should have been the zero value")                                //"Expected: (zero value) Actual: 123")
	fixture.fail(so(" ", ShouldBeZeroValue), "| |' ' should have been the zero value")                                      //"Expected: (zero value) Actual:  ")
	fixture.fail(so([]string{"Nonempty"}, ShouldBeZeroValue), "[]|[Nonempty]|'[Nonempty]' should have been the zero value") //"Expected: (zero value) Actual: [Nonempty]")
	fixture.fail(so(struct{ a string }{a: "asdf"}, ShouldBeZeroValue), "{}|{asdf}|'{a:asdf}' should have been the zero value")
	fixture.pass(so(0, ShouldBeZeroValue))
	fixture.pass(so(false, ShouldBeZeroValue))
	fixture.pass(so("", ShouldBeZeroValue))
	fixture.pass(so(struct{}{}, ShouldBeZeroValue))
}
