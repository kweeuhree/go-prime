package main

import "testing"

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
