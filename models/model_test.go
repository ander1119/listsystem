package models

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestPageInputValidation(t *testing.T) {
	// valid input
	validInput := PageInput{
		Articles: "Lorem ipsum dolor sit amet",
	}
	validate := validator.New()

	err := validate.Struct(validInput)
	if err != nil {
		t.Errorf("validation failed for valid input: %v", err)
	}

	// invalid input
	invalidInput := PageInput{}

	err = validate.Struct(invalidInput)
	if err == nil {
		t.Errorf("validation passed for invalid input")
	}

	// another invalid input
	invalidInput2 := PageInput{
		Articles:    "",
		NextPageKey: "",
	}

	err = validate.Struct(invalidInput2)
	if err == nil {
		t.Errorf("validation passed for invalid input")
	}
}

func TestHeadInputValidation(t *testing.T) {
	// valid input
	validInput := HeadInput{
		NextPageKey: "test_next_page_key",
	}
	validate := validator.New()

	err := validate.Struct(validInput)
	if err != nil {
		t.Errorf("validation failed for valid input: %v", err)
	}

	// invalid input
	invalidInput := HeadInput{}

	err = validate.Struct(invalidInput)
	if err == nil {
		t.Errorf("validation passed for invalid input")
	}

	// another invalid input
	invalidInput2 := HeadInput{
		NextPageKey: "",
	}

	err = validate.Struct(invalidInput2)
	if err == nil {
		t.Errorf("validation passed for invalid input")
	}
}