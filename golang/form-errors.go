// ========== ./internal/forms/errors.go
package forms

type errors map[string][]string

// Add adds error message to given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get return first error message
func (e errors) Get(field string) string {
	if len(e[field]) == 0 {
		return ""
	}

	return e[field][0]
}

// ========== ./internal/forms/forms.go
package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form stores form values and errors
type Form struct {
	Values url.Values
	Errors errors
}

// Valid return true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initilizes Form data
func New(data url.Values) *Form {
	return &Form{
		data,
		//errors(map[string][]string{}),
		errors{},
	}
}

// Has checks if request has field
func (f *Form) Has(field string) bool {
	value := f.Values.Get(field)
	if value == "" {
		f.Errors.Add(field, fmt.Sprintf("%s can't be blank", field))
		return false
	}

	return true
}

// Required checks all given fields have value
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, fmt.Sprintf("%s can't be blank", field))
		}
	}
}

// MinLength checks field lenght not less then given length
func (f *Form) MinLength(field string, length int) {
	value := f.Values.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("%s can't be less than %d symbols", field, length))
	}
}

// IsEmail checks email valid
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Values.Get(field)) {
		f.Errors.Add(field, fmt.Sprintf("%s is not valid", field))
	}
}


