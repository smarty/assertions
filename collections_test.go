package assertions

import (
	"fmt"
	"time"
)

func (fixture *AssertionsFixture) TestShouldContainKey() {
	fixture.fail(so(map[int]int{}, ShouldContainKey), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(map[int]int{}, ShouldContainKey, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(Thing1{}, ShouldContainKey, 1), "You must provide a valid map type (was assertions.Thing1)!")
	fixture.fail(so(nil, ShouldContainKey, 1), "You must provide a valid map type (was <nil>)!")
	fixture.fail(so(map[int]int{1: 41}, ShouldContainKey, 2), "Expected the map[int]int to contain the key: [2] (but it didn't)!")

	fixture.pass(so(map[int]int{1: 41}, ShouldContainKey, 1))
	fixture.pass(so(map[int]int{1: 41, 2: 42, 3: 43}, ShouldContainKey, 2))
}

func (fixture *AssertionsFixture) TestShouldNotContainKey() {
	fixture.fail(so(map[int]int{}, ShouldNotContainKey), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so(map[int]int{}, ShouldNotContainKey, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(Thing1{}, ShouldNotContainKey, 1), "You must provide a valid map type (was assertions.Thing1)!")
	fixture.fail(so(nil, ShouldNotContainKey, 1), "You must provide a valid map type (was <nil>)!")
	fixture.fail(so(map[int]int{1: 41}, ShouldNotContainKey, 1), "Expected the map[int]int NOT to contain the key: [1] (but it did)!")
	fixture.pass(so(map[int]int{1: 41}, ShouldNotContainKey, 2))
}

func (fixture *AssertionsFixture) TestShouldContain() {
	fixture.fail(so([]int{}, ShouldContain), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so([]int{}, ShouldContain, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(Thing1{}, ShouldContain, 1), "You must provide a valid container (was assertions.Thing1)!")
	fixture.fail(so(nil, ShouldContain, 1), "You must provide a valid container (was <nil>)!")
	fixture.fail(so([]int{1}, ShouldContain, 2), "Expected the container ([]int) to contain: '2' (but it didn't)!")
	fixture.fail(so([][]int{{1}}, ShouldContain, []int{2}), "Expected the container ([][]int) to contain: '[2]' (but it didn't)!")

	fixture.pass(so([]int{1}, ShouldContain, 1))
	fixture.pass(so([]int{1, 2, 3}, ShouldContain, 2))
	fixture.pass(so([][]int{{1}, {2}, {3}}, ShouldContain, []int{2}))
}

func (fixture *AssertionsFixture) TestShouldNotContain() {
	fixture.fail(so([]int{}, ShouldNotContain), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so([]int{}, ShouldNotContain, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fixture.fail(so(Thing1{}, ShouldNotContain, 1), "You must provide a valid container (was assertions.Thing1)!")
	fixture.fail(so(nil, ShouldNotContain, 1), "You must provide a valid container (was <nil>)!")

	fixture.fail(so([]int{1}, ShouldNotContain, 1), "Expected the container ([]int) NOT to contain: '1' (but it did)!")
	fixture.fail(so([]int{1, 2, 3}, ShouldNotContain, 2), "Expected the container ([]int) NOT to contain: '2' (but it did)!")
	fixture.fail(so([][]int{{1}, {2}, {3}}, ShouldNotContain, []int{2}), "Expected the container ([][]int) NOT to contain: '[2]' (but it did)!")

	fixture.pass(so([]int{1}, ShouldNotContain, 2))
	fixture.pass(so([][]int{{1}, {2}, {3}}, ShouldNotContain, []int{4}))
}

func (fixture *AssertionsFixture) TestShouldBeIn() {
	fixture.fail(so(4, ShouldBeIn), needNonEmptyCollection)

	container := []int{1, 2, 3, 4}
	fixture.pass(so(4, ShouldBeIn, container))
	fixture.pass(so(4, ShouldBeIn, 1, 2, 3, 4))
	fixture.pass(so([]int{4}, ShouldBeIn, [][]int{{1}, {2}, {3}, {4}}))
	fixture.pass(so([]int{4}, ShouldBeIn, []int{1}, []int{2}, []int{3}, []int{4}))

	fixture.fail(so(4, ShouldBeIn, 1, 2, 3), "Expected '4' to be in the container ([]interface {}), but it wasn't!")
	fixture.fail(so(4, ShouldBeIn, []int{1, 2, 3}), "Expected '4' to be in the container ([]int), but it wasn't!")
	fixture.fail(so([]int{4}, ShouldBeIn, []int{1}, []int{2}, []int{3}), "Expected '[4]' to be in the container ([]interface {}), but it wasn't!")
	fixture.fail(so([]int{4}, ShouldBeIn, [][]int{{1}, {2}, {3}}), "Expected '[4]' to be in the container ([][]int), but it wasn't!")
}

func (fixture *AssertionsFixture) TestShouldNotBeIn() {
	fixture.fail(so(4, ShouldNotBeIn), needNonEmptyCollection)

	container := []int{1, 2, 3, 4}
	fixture.pass(so(42, ShouldNotBeIn, container))
	fixture.pass(so(42, ShouldNotBeIn, 1, 2, 3, 4))
	fixture.pass(so([]int{42}, ShouldNotBeIn, []int{1}, []int{2}, []int{3}, []int{4}))
	fixture.pass(so([]int{42}, ShouldNotBeIn, [][]int{{1}, {2}, {3}, {4}}))

	fixture.fail(so(2, ShouldNotBeIn, 1, 2, 3), "Expected '2' NOT to be in the container ([]interface {}), but it was!")
	fixture.fail(so(2, ShouldNotBeIn, []int{1, 2, 3}), "Expected '2' NOT to be in the container ([]int), but it was!")
	fixture.fail(so([]int{2}, ShouldNotBeIn, []int{1}, []int{2}, []int{3}), "Expected '[2]' NOT to be in the container ([]interface {}), but it was!")
	fixture.fail(so([]int{2}, ShouldNotBeIn, [][]int{{1}, {2}, {3}}), "Expected '[2]' NOT to be in the container ([][]int), but it was!")
}

func (fixture *AssertionsFixture) TestShouldBeEmpty() {
	fixture.fail(so(1, ShouldBeEmpty, 2, 3), "This assertion requires exactly 0 comparison values (you provided 2).")

	fixture.pass(so([]int{}, ShouldBeEmpty))           // empty slice
	fixture.pass(so([][]int{}, ShouldBeEmpty))         // empty slice
	fixture.pass(so([]interface{}{}, ShouldBeEmpty))   // empty slice
	fixture.pass(so(map[string]int{}, ShouldBeEmpty))  // empty map
	fixture.pass(so("", ShouldBeEmpty))                // empty string
	fixture.pass(so(&[]int{}, ShouldBeEmpty))          // pointer to empty slice
	fixture.pass(so(&[0]int{}, ShouldBeEmpty))         // pointer to empty array
	fixture.pass(so(nil, ShouldBeEmpty))               // nil
	fixture.pass(so(make(chan string), ShouldBeEmpty)) // empty channel

	fixture.fail(so([]int{1}, ShouldBeEmpty), "Expected [1] to be empty (but it wasn't)!")                      // non-empty slice
	fixture.fail(so([][]int{{1}}, ShouldBeEmpty), "Expected [[1]] to be empty (but it wasn't)!")                // non-empty slice
	fixture.fail(so([]interface{}{1}, ShouldBeEmpty), "Expected [1] to be empty (but it wasn't)!")              // non-empty slice
	fixture.fail(so(map[string]int{"hi": 0}, ShouldBeEmpty), "Expected map[hi:0] to be empty (but it wasn't)!") // non-empty map
	fixture.fail(so("hi", ShouldBeEmpty), "Expected hi to be empty (but it wasn't)!")                           // non-empty string
	fixture.fail(so(&[]int{1}, ShouldBeEmpty), "Expected &[1] to be empty (but it wasn't)!")                    // pointer to non-empty slice
	fixture.fail(so(&[1]int{1}, ShouldBeEmpty), "Expected &[1] to be empty (but it wasn't)!")                   // pointer to non-empty array
	c := make(chan int, 1)                                                                                      // non-empty channel
	go func() { c <- 1 }()
	time.Sleep(time.Millisecond)
	fixture.fail(so(c, ShouldBeEmpty), fmt.Sprintf("Expected %+v to be empty (but it wasn't)!", c))
}

func (fixture *AssertionsFixture) TestShouldNotBeEmpty() {
	fixture.fail(so(1, ShouldNotBeEmpty, 2, 3), "This assertion requires exactly 0 comparison values (you provided 2).")

	fixture.fail(so([]int{}, ShouldNotBeEmpty), "Expected [] to NOT be empty (but it was)!")             // empty slice
	fixture.fail(so([]interface{}{}, ShouldNotBeEmpty), "Expected [] to NOT be empty (but it was)!")     // empty slice
	fixture.fail(so(map[string]int{}, ShouldNotBeEmpty), "Expected map[] to NOT be empty (but it was)!") // empty map
	fixture.fail(so("", ShouldNotBeEmpty), "Expected  to NOT be empty (but it was)!")                    // empty string
	fixture.fail(so(&[]int{}, ShouldNotBeEmpty), "Expected &[] to NOT be empty (but it was)!")           // pointer to empty slice
	fixture.fail(so(&[0]int{}, ShouldNotBeEmpty), "Expected &[] to NOT be empty (but it was)!")          // pointer to empty array
	fixture.fail(so(nil, ShouldNotBeEmpty), "Expected <nil> to NOT be empty (but it was)!")              // nil
	c := make(chan int)                                                                                  // non-empty channel
	fixture.fail(so(c, ShouldNotBeEmpty), fmt.Sprintf("Expected %+v to NOT be empty (but it was)!", c))  // empty channel

	fixture.pass(so([]int{1}, ShouldNotBeEmpty))                // non-empty slice
	fixture.pass(so([]interface{}{1}, ShouldNotBeEmpty))        // non-empty slice
	fixture.pass(so(map[string]int{"hi": 0}, ShouldNotBeEmpty)) // non-empty map
	fixture.pass(so("hi", ShouldNotBeEmpty))                    // non-empty string
	fixture.pass(so(&[]int{1}, ShouldNotBeEmpty))               // pointer to non-empty slice
	fixture.pass(so(&[1]int{1}, ShouldNotBeEmpty))              // pointer to non-empty array
	c = make(chan int, 1)
	go func() { c <- 1 }()
	time.Sleep(time.Millisecond)
	fixture.pass(so(c, ShouldNotBeEmpty))
}

func (fixture *AssertionsFixture) TestShouldHaveLength() {
	fixture.fail(so(1, ShouldHaveLength, 2), "You must provide a valid container (was int)!")
	fixture.fail(so(nil, ShouldHaveLength, 1), "You must provide a valid container (was <nil>)!")
	fixture.fail(so("hi", ShouldHaveLength, float64(1.0)), "You must provide a valid integer (was float64)!")
	fixture.fail(so([]string{}, ShouldHaveLength), "This assertion requires exactly 1 comparison values (you provided 0).")
	fixture.fail(so([]string{}, ShouldHaveLength, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	fixture.fail(so([]string{}, ShouldHaveLength, -10), "You must provide a valid positive integer (was -10)!")

	fixture.fail(so([]int{}, ShouldHaveLength, 1), "Expected [] (length: 0) to have length equal to '1', but it wasn't!")             // empty slice
	fixture.fail(so([]interface{}{}, ShouldHaveLength, 1), "Expected [] (length: 0) to have length equal to '1', but it wasn't!")     // empty slice
	fixture.fail(so(map[string]int{}, ShouldHaveLength, 1), "Expected map[] (length: 0) to have length equal to '1', but it wasn't!") // empty map
	fixture.fail(so("", ShouldHaveLength, 1), "Expected  (length: 0) to have length equal to '1', but it wasn't!")                    // empty string
	fixture.fail(so(&[]int{}, ShouldHaveLength, 1), "Expected &[] (length: 0) to have length equal to '1', but it wasn't!")           // pointer to empty slice
	fixture.fail(so(&[0]int{}, ShouldHaveLength, 1), "Expected &[] (length: 0) to have length equal to '1', but it wasn't!")          // pointer to empty array
	c := make(chan int)                                                                                                               // non-empty channel
	fixture.fail(so(c, ShouldHaveLength, 1), fmt.Sprintf("Expected %+v (length: 0) to have length equal to '1', but it wasn't!", c))
	c = make(chan int) // empty channel
	fixture.fail(so(c, ShouldHaveLength, 1), fmt.Sprintf("Expected %+v (length: 0) to have length equal to '1', but it wasn't!", c))

	fixture.pass(so([]int{1}, ShouldHaveLength, 1))                // non-empty slice
	fixture.pass(so([]interface{}{1}, ShouldHaveLength, 1))        // non-empty slice
	fixture.pass(so(map[string]int{"hi": 0}, ShouldHaveLength, 1)) // non-empty map
	fixture.pass(so("hi", ShouldHaveLength, 2))                    // non-empty string
	fixture.pass(so(&[]int{1}, ShouldHaveLength, 1))               // pointer to non-empty slice
	fixture.pass(so(&[1]int{1}, ShouldHaveLength, 1))              // pointer to non-empty array
	c = make(chan int, 1)
	go func() { c <- 1 }()
	time.Sleep(time.Millisecond)
	fixture.pass(so(c, ShouldHaveLength, 1))
	fixture.pass(so(c, ShouldHaveLength, uint(1)))
}
