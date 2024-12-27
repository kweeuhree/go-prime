package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	// create table struct
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		// define entries
		{"prime", 2, true, "2 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers aren't prime, by definition!"},
	}

	// loop through primeTests
	for _, entry := range primeTests {
		result, msg := isPrime(entry.testNum)

		if entry.expected && !result {
			t.Errorf("%s: expected true, but got false", entry.name)
		}

		if !entry.expected && result {
			t.Errorf("%s: expected false, but got true", entry.name)
		}

		if entry.msg != msg {
			t.Errorf("%s: expected %s, but got %s", entry.name, entry.msg, msg)
		}
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	currentOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to write pipe
	os.Stdout = w

	intro()

	// close writer in order to avoid a resource leak
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = currentOut

	// read the output of intro function from the read pipe
	// this returns a slice of bites
	out, _ := io.ReadAll(r)

	// perform the test
	// cast the slice of bites to a string
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("incorrect intro: expected -. but got %s", string(out))
	}

}
