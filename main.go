package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()
	// create a channel to indicate when a user wants to quit
	doneChannel := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(doneChannel)
	// block until the doneChannel gets a value
	<-doneChannel
	// close the channel
	close(doneChannel)
	// say goodbuy
	fmt.Println("Goodbye.")
}

func readUserInput(doneChannel chan bool) {
	// read user input
	scanner := bufio.NewScanner(os.Stdin)

	// run until its done
	for {
		result, done := checkNumbers(scanner)

		if done {
			doneChannel <- true
			return
		}

		fmt.Println(result)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert user input into an int
	numberToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number.", false
	}

	_, msg := isPrime(numberToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number to find out if the number is prime. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers aren't prime
	if n < 0 {
		return false, "Negative numbers aren't prime, by definition!"
	}

	// use modulus operator to check for prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
