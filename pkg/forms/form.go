package forms

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Form struct {
	Errors errors
	url.Values
}

func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: make(errors),
	}
}

// Goes for every field in the slice and:
// checks that the field values's length is not 0,
// if so add an error message.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Checks the field value for any whitespace character,
// if it contains whitespace, add an error message for the field.
func (f *Form) NoWhiteSpace(field string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	for _, c := range value {
		if unicode.IsSpace(c) {
			f.Errors.Add(field, "This field cannot contain whitespace")
			return
		}
	}
}

// Ensures the field value has a minimum of minLength characters.
func (f *Form) MinLength(field string, minLength int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) < minLength {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum length is %d)", minLength))
	}
}

// Checks whether the fields values are of type int,
// this is done by trying to convert the value
// of the field to an int using the Atoi function.
func (f *Form) RequireTypeInt(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if value == "" {
			return
		}

		_, err := strconv.Atoi(value)
		if err != nil {
			f.Errors.Add(field, "Please enter a valid number")
		}
	}
}

// Ensures the field value has a maximum of maxLength characters.
func (f *Form) MaxLength(field string, maxLength int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > maxLength {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum length is %d)", maxLength))
	}
}

// Returns true if there's zero errors.
func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}
