package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername  = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFirstName = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	isValidLastName  = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 5, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letter, digits or underscore")
	}
	return nil
}

func ValidateFirstName(value string) error {
	if err := ValidateString(value, 2, 100); err != nil {
		return err
	}
	if !isValidFirstName(value) {
		return fmt.Errorf("must contain only letter")
	}
	return nil
}

func ValidateLastName(value string) error {
	if err := ValidateString(value, 2, 100); err != nil {
		return err
	}
	if !isValidLastName(value) {
		return fmt.Errorf("must contain only letter")
	}
	return nil
}

func ValidatePassword(value string) error {
	if err := ValidateString(value, 8, 100); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("%s is not a valid email address", value)
	}
	return nil
}

func ValidateAdress(value string) error {
	if err := ValidateString(value, 5, 100); err != nil {
		return err
	}
	return nil
}

func ValidateProductName(value string) error {
	if err := ValidateString(value, 2, 100); err != nil {
		return err
	}
	return nil
}

func ValidateProductPrice(value int) error {
	if value < 0 {
		return fmt.Errorf("product price cannot below 0")
	}
	return nil
}

func ValidateProductQuantity(value int) error {
	if value < 0 {
		return fmt.Errorf("product product cannot below 0")
	}
	return nil
}
