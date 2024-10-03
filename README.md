# Persian Phone Number Validator

A simple Go package to clean, normalize, and validate Persian (Iranian) phone numbers. The package supports various phone number formats, including international ones (e.g., `+98`, `0098`) and local numbers starting with `09`.

## Features

- **Clean Phone Numbers**: Removes all non-numeric characters like spaces, parentheses, dashes, etc.
- **Normalize Phone Numbers**: Converts numbers with country code (`+98`, `0098`) to local format (`09xxxxxxxxx`).
- **Validate Phone Numbers**: Ensures phone numbers are in the correct format (11 digits, starting with `09`).
- **Combined Functionality**: Normalize and validate in a single function call.

## Installation

To install the package, use:

```bash
go get github.com/yourusername/persianPhoneValidator
```

## Usage

### Example Code

```go
package main

import (
	"fmt"
	"persianPhoneValidator"
)

func main() {
	phone := "+989121234567"

	// Clean the phone number
	cleanedPhone := persianPhoneValidator.CleanPhoneNumber(phone)
	fmt.Println("Cleaned Phone:", cleanedPhone)

	// Normalize the phone number
	normalizedPhone := persianPhoneValidator.NormalizePhoneNumber(phone)
	fmt.Println("Normalized Phone:", normalizedPhone)

	// Check if the phone number is valid
	isValid := persianPhoneValidator.IsPhoneNumberValid(phone)
	fmt.Println("Is Phone Valid:", isValid)

	// Validate and normalize the phone number
	isValid, normalized := persianPhoneValidator.ValidatePhoneNumber(phone)
	fmt.Printf("Phone: %s, Valid: %v\n", normalized, isValid)
}
```

### Function Descriptions

#### `CleanPhoneNumber(phone string) string`

Removes all non-numeric characters from the phone number.

**Example:**

```go
fmt.Println(persianPhoneValidator.CleanPhoneNumber("0912-123-4567"))
// Output: "09121234567"
```

#### `NormalizePhoneNumber(phone string) string`

Converts a phone number to the local format (`09xxxxxxxxx`) by handling different international formats (`+98`, `0098`).

**Example:**

```go
fmt.Println(persianPhoneValidator.NormalizePhoneNumber("+989121234567"))
// Output: "09121234567"
```

#### `IsPhoneNumberValid(phone string) bool`

Validates if the phone number is valid. A valid number has 11 digits and starts with `09`.

**Example:**

```go
fmt.Println(persianPhoneValidator.IsPhoneNumberValid("09121234567"))
// Output: true
```

#### `ValidatePhoneNumber(phone string) (bool, string)`

Normalizes and validates the phone number. Returns `true` and the normalized number if valid; otherwise, returns `false` and the normalized number.

**Example:**

```go
isValid, normalizedPhone := persianPhoneValidator.ValidatePhoneNumber("+989121234567")
fmt.Println("Valid:", isValid, "Normalized:", normalizedPhone)
// Output: Valid: true, Normalized: 09121234567
```

## Running Tests

To run the unit tests, execute the following command in the root of your project:

```bash
go test .
```

This will run the comprehensive test suite for the package, ensuring all functions work as expected.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contribution

Feel free to open issues or submit pull requests. Contributions are always welcome!
