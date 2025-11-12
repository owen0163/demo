package greet_test

import (
	"testing"
)

func TestGreetNmae(t *testing.T) {
	given := "owen"
	expact := "Hi, owen."

	actual := greet(given)

	if expact != actual {
		t.Errorf("given %q ,  expact %q, actual %q\n", given, actual, expact)
	}

}
