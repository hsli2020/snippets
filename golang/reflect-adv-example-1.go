package main

import (
  "fmt"
  "reflect"
  "regexp"
  "strings"
)

// Name of the struct tag used in examples.
const tagName = "validate"

// Regular expression to validate email address.
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// Generic data validator.
type Validator interface {
  // Validate method performs validation and returns result and optional error.
  Validate(interface{}) (bool, error)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
  return true, nil
}

// StringValidator validates string presence and/or its length.
type StringValidator struct {
  Min int
  Max int
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
  l := len(val.(string))

  if l == 0 {
    return false, fmt.Errorf("cannot be blank")
  }

  if l < v.Min {
    return false, fmt.Errorf("should be at least %v chars long", v.Min)
  }

  if v.Max >= v.Min && l > v.Max {
    return false, fmt.Errorf("should be less than %v chars long", v.Max)
  }

  return true, nil
}

// NumberValidator performs numerical value validation.
// Its limited to int type for simplicity.
type NumberValidator struct {
  Min int
  Max int
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
  num := val.(int)

  if num < v.Min {
    return false, fmt.Errorf("should be greater than %v", v.Min)
  }

  if v.Max >= v.Min && num > v.Max {
    return false, fmt.Errorf("should be less than %v", v.Max)
  }

  return true, nil
}

// EmailValidator checks if string is a valid email address.
type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
  if !mailRe.MatchString(val.(string)) {
    return false, fmt.Errorf("is not a valid email address")
  }
  return true, nil
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
  args := strings.Split(tag, ",")

  switch args[0] {
  case "number":
    validator := NumberValidator{}
    fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
    return validator
  case "string":
    validator := StringValidator{}
    fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
    return validator
  case "email":
    return EmailValidator{}
  }

  return DefaultValidator{}
}

// Performs actual data validation using validator definitions on the struct
func validateStruct(s interface{}) []error {
  errs := []error{}

  // ValueOf returns a Value representing the run-time data
  v := reflect.ValueOf(s)

  for i := 0; i < v.NumField(); i++ {
    // Get the field tag value
    tag := v.Type().Field(i).Tag.Get(tagName)

    // Skip if tag is not defined or ignored
    if tag == "" || tag == "-" {
      continue
    }

    // Get a validator that corresponds to a tag
    validator := getValidatorFromTag(tag)

    // Perform validation
    valid, err := validator.Validate(v.Field(i).Interface())

    // Append error to results
    if !valid && err != nil {
      errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
    }
  }

  return errs
}

type User struct {
  Id    int    `validate:"number,min=1,max=1000"`
  Name  string `validate:"string,min=2,max=10"`
  Bio   string `validate:"string"`
  Email string `validate:"email"`
}

func main() {
  user := User{
    Id:    0,
    Name:  "superlongstring",
    Bio:   "",
    Email: "foobar",
  }

  fmt.Println("Errors:")
  for i, err := range validateStruct(user) {
    fmt.Printf("\t%d. %s\n", i+1, err.Error())
  }
}