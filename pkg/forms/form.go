package forms

import (
	"fmt"
	"net/url"
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

// Returns true if there's zero errors.
func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}
