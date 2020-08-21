// https://github.com/golang/go/issues/52607
// https://go.dev/play/p/usV6mFRjxxb

package main

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

// With returns an error that represents the union of the two errors.
// It will report true from errors.Is and errors.As if calling
// that function would return true on either err or other.
// Calling errors.Unwrap on the error will return err.
func With(err, other error) error {
	if err == nil {
		return nil
	}
	return union{error: err, other: other}
}

type union struct {
	error
	other error
}

func (u union) Is(target error) bool {
	if errors.Is(u.other, target) {
		return true
	}
	return errors.Is(u.error, target)
}

func (u union) As(target any) bool {
	if errors.As(u.other, target) {
		return true
	}
	return errors.As(u.error, target)
}

func (u union) Unwrap() error {
	return u.error
}

// Tests Below

var NotFound = errors.New("not found")

func TestWithIs(t *testing.T) {
	err := errors.New("some pig")
	wrapped := fmt.Errorf("wilbur: %w", err)
	err2 := With(wrapped, NotFound)

	if !errors.Is(err2, NotFound) {
		t.Fatal("failed to find flag")
	}
	if !errors.Is(err2, err) {
		t.Fatal("failed to find original error")
	}
	if !errors.Is(err2, wrapped) {
		t.Fatal("failed to find wrapped error")
	}

	err3 := fmt.Errorf("more context: %w", With(err2, io.EOF))

	if !errors.Is(err3, NotFound) {
		t.Fatal("failed to find flag after wrapping")
	}
	if !errors.Is(err3, err) {
		t.Fatal("failed to find original error after wrapping")
	}
	if !errors.Is(err3, wrapped) {
		t.Fatal("failed to find wrapped error after second wrapping")
	}
	if !errors.Is(err3, io.EOF) {
		t.Fatal("failed to find flagged wrapped error after wrapping")
	}
}

type myError string

func (m myError) Error() string {
	return string(m)
}

type otherError struct {
	msg string
}

func (o otherError) Error() string {
	return o.msg
}

func TestWithAs(t *testing.T) {
	err := myError("some pig")
	wrapped := fmt.Errorf("wilbur: %w", err)
	err2 := With(wrapped, NotFound)
	err3 := fmt.Errorf("more context: %w", err2)

	var my myError

	if !errors.As(err3, &my) {
		t.Fatal("failed to original type after wrapping")
	}

	other := otherError{msg: "hi!"}

	err4 := With(err3, fmt.Errorf("some other error: %w", other))

	var o otherError
	if !errors.As(err4, &o) {
		t.Fatal("failed to find flagged type")
	}
	if !errors.As(err4, &my) {
		t.Fatal("failed to original type after wrapping")
	}

	if !errors.Is(err4, err) {
		t.Fatal("failed to find original error")
	}
	if !errors.Is(err4, other) {
		t.Fatal("failed to find flagged error")
	}
}
