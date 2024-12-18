# **Prime Number Checker CLI**

## **Overview**

The **Prime Number Checker CLI** is a command-line application designed to determine if a given number is prime. Users can interact with the application by entering numbers via the terminal. The app provides immediate feedback and supports quitting the session with the command `q`.

This project aims to serve as a foundation for a comprehensive testing suite, targeting **90%+ test coverage** to ensure robustness and reliability.

---

## **Features**

- Accepts integer input and determines if the number is prime.
- Handles invalid input gracefully, providing user-friendly feedback.
- Quits the program cleanly when the user enters `q`.
- Includes modular code to facilitate test coverage and future extensions.

---

## **Installation**

1. **Clone the Repository**:

   ```bash
   git clone <repository-url>
   cd prime-number-checker

   ```

2. **Run the Program: Build and execute the application:**:

   ```bash
   go run main.go
   ```

3. **Run Tests: Execute the tests to verify the implementation:**:

   ```bash
   go test .

   ```

## **Usage**

1. **Start the application**:

   ```bash
   go run main.go
   ```

2. **Follow the prompts**:
   - Enter a number to check if it is prime.
   - Enter `q` to quit the application.

**Example Interaction**:

```
Is it Prime?
------------
Enter a whole number to find out if the number is prime. Enter q to quit.
-> 7
7 is a prime number!
-> 8
8 is not a prime number because it is divisible by 2
-> -5
Negative numbers aren't prime, by definition!
-> q
Goodbye.

```

## **Code Overview**

**Main Components**
| Function | Description |
|-------------------|---------------------------------------------------------------------------|
| `main()` | Initializes the application and manages the main event loop. |
| `intro()` | Prints a welcome message and usage instructions. |
| `readUserInput()` | Continuously listens for user input, processes it, and exits when `q` is entered. |
| `checkNumbers()` | Validates user input and determines if the number is prime. |
| `isPrime()` | Core logic for checking if a number is prime. |

## **Testing**

The application includes unit tests for the `isPrime` function, ensuring correct results across various edge cases and inputs.

### **Test Cases**

| Test Name          | Input | Expected Output                                      |
| ------------------ | ----- | ---------------------------------------------------- |
| Prime Number       | 2     | 2 is a prime number!                                 |
| Not a Prime Number | 8     | 8 is not a prime number because it is divisible by 2 |
| Zero               | 0     | 0 is not prime, by definition!                       |
| Negative Number    | -11   | Negative numbers aren't prime, by definition!        |

### **Planned Enhancements**

Expand test coverage to include additional functions and edge cases.
