package assertions

func (fixture *AssertionsFixture) TestShouldStartWith() {
	fixture.fail(so("", ShouldStartWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so("", ShouldStartWith, "asdf", "asdf"), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.pass(so("", ShouldStartWith, ""))
	fixture.fail(so("", ShouldStartWith, "x"), "x||Expected '' to start with 'x' (but it didn't)!")
	fixture.pass(so("abc", ShouldStartWith, "abc"))
	fixture.fail(so("abc", ShouldStartWith, "abcd"), "abcd|abc|Expected 'abc' to start with 'abcd' (but it didn't)!")

	fixture.pass(so("superman", ShouldStartWith, "super"))
	fixture.fail(so("superman", ShouldStartWith, "bat"), "bat|sup...|Expected 'superman' to start with 'bat' (but it didn't)!")
	fixture.fail(so("superman", ShouldStartWith, "man"), "man|sup...|Expected 'superman' to start with 'man' (but it didn't)!")

	fixture.fail(so(1, ShouldStartWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func (fixture *AssertionsFixture) TestShouldNotStartWith() {
	fixture.fail(so("", ShouldNotStartWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so("", ShouldNotStartWith, "asdf", "asdf"), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.fail(so("", ShouldNotStartWith, ""), "Expected '<empty>' NOT to start with '<empty>' (but it did)!")
	fixture.fail(so("superman", ShouldNotStartWith, "super"), "Expected 'superman' NOT to start with 'super' (but it did)!")
	fixture.pass(so("superman", ShouldNotStartWith, "bat"))
	fixture.pass(so("superman", ShouldNotStartWith, "man"))

	fixture.fail(so(1, ShouldNotStartWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func (fixture *AssertionsFixture) TestShouldEndWith() {
	fixture.fail(so("", ShouldEndWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so("", ShouldEndWith, "", ""), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.pass(so("", ShouldEndWith, ""))
	fixture.fail(so("", ShouldEndWith, "z"), "z||Expected '' to end with 'z' (but it didn't)!")
	fixture.pass(so("xyz", ShouldEndWith, "xyz"))
	fixture.fail(so("xyz", ShouldEndWith, "wxyz"), "wxyz|xyz|Expected 'xyz' to end with 'wxyz' (but it didn't)!")

	fixture.pass(so("superman", ShouldEndWith, "man"))
	fixture.fail(so("superman", ShouldEndWith, "super"), "super|...erman|Expected 'superman' to end with 'super' (but it didn't)!")
	fixture.fail(so("superman", ShouldEndWith, "blah"), "blah|...rman|Expected 'superman' to end with 'blah' (but it didn't)!")

	fixture.fail(so(1, ShouldEndWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func (fixture *AssertionsFixture) TestShouldNotEndWith() {
	fixture.fail(so("", ShouldNotEndWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so("", ShouldNotEndWith, "", ""), "This assertion requires exactly 1 comparison values (you provided 2).")

	fixture.fail(so("", ShouldNotEndWith, ""), "Expected '<empty>' NOT to end with '<empty>' (but it did)!")
	fixture.fail(so("superman", ShouldNotEndWith, "man"), "Expected 'superman' NOT to end with 'man' (but it did)!")
	fixture.pass(so("superman", ShouldNotEndWith, "super"))

	fixture.fail(so(1, ShouldNotEndWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func (fixture *AssertionsFixture) TestShouldContainSubstring() {
	fixture.fail(so("asdf", ShouldContainSubstring), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so("asdf", ShouldContainSubstring, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(123, ShouldContainSubstring, 23), "Both arguments to this assertion must be strings (you provided int and int).")

	fixture.pass(so("asdf", ShouldContainSubstring, "sd"))
	fixture.fail(so("qwer", ShouldContainSubstring, "sd"), "sd|qwer|Expected 'qwer' to contain substring 'sd' (but it didn't)!")
}

func (fixture *AssertionsFixture) TestShouldNotContainSubstring() {
	fixture.fail(so("asdf", ShouldNotContainSubstring), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so("asdf", ShouldNotContainSubstring, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(123, ShouldNotContainSubstring, 23), "Both arguments to this assertion must be strings (you provided int and int).")

	fixture.pass(so("qwer", ShouldNotContainSubstring, "sd"))
	fixture.fail(so("asdf", ShouldNotContainSubstring, "sd"), "Expected 'asdf' NOT to contain substring 'sd' (but it did)!")
}

func (fixture *AssertionsFixture) TestShouldBeBlank() {
	fixture.fail(so("", ShouldBeBlank, "adsf"), "This assertion requires exactly 0 comparison values (you provided 1).")
	fixture.fail(so(1, ShouldBeBlank), "The argument to this assertion must be a string (you provided int).")

	fixture.fail(so("asdf", ShouldBeBlank), "|asdf|Expected 'asdf' to be blank (but it wasn't)!")
	fixture.pass(so("", ShouldBeBlank))
}

func (fixture *AssertionsFixture) TestShouldNotBeBlank() {
	fixture.fail(so("", ShouldNotBeBlank, "adsf"), "This assertion requires exactly 0 comparison values (you provided 1).")
	fixture.fail(so(1, ShouldNotBeBlank), "The argument to this assertion must be a string (you provided int).")

	fixture.fail(so("", ShouldNotBeBlank), "Expected value to NOT be blank (but it was)!")
	fixture.pass(so("asdf", ShouldNotBeBlank))
}

func (fixture *AssertionsFixture) TestShouldEqualWithout() {
	fixture.fail(so("", ShouldEqualWithout, ""), "This assertion requires exactly 2 comparison values (you provided 1).")
	fixture.fail(so(1, ShouldEqualWithout, 2, 3), "All arguments to this assertion must be strings (you provided: [int int int]).")

	fixture.fail(so("asdf", ShouldEqualWithout, "qwer", "q"), "Expected 'asdf' to equal 'qwer' but without any 'q' (but it didn't).")
	fixture.pass(so("asdf", ShouldEqualWithout, "df", "as"))
}

func (fixture *AssertionsFixture) TestShouldEqualTrimSpace() {
	fixture.fail(so(" asdf ", ShouldEqualTrimSpace), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(1, ShouldEqualTrimSpace, 2), "Both arguments to this assertion must be strings (you provided int and int).")

	fixture.fail(so("asdf", ShouldEqualTrimSpace, "qwer"), "qwer|asdf|Expected: 'qwer' Actual: 'asdf' (Should be equal)")
	fixture.pass(so(" asdf\t\n", ShouldEqualTrimSpace, "asdf"))
}
