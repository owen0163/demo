package greet

import (
	"testing"
)

func TestGreetName(t *testing.T) {
	given := "owen"
	expect := "Hi, owen."

	actual := greet(given)

	if expect != actual {
		t.Errorf("given %q ,  expect %q, actual %q\n", given, actual, expect)
	}

}
