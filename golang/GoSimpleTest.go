// https://twitter.com/julienBisconti/status/1556024007434149889

package main

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestExample(t *testing.T) {
	type TT struct {
		got  string
		want string
	}

	assert := func(t *testing.T, tt TT) {
		if diff := Diff(tt.got, tt.want); diff != "" {
			t.Error(Callers(), diff)
		}
	}

	t.Run("1", func(t *testing.T) {
		t.Parallel()
		assert(t, TT{"lorem ipsum dolor amet", "lorem ipsum dolor sit amet"})
	})

	t.Run("2", func(t *testing.T) {
		t.Parallel()
		assert(t, TT{"the quick fox jumped over lazy dog", "the quick brown fox jumped over the lazy dog"})
	})

	t.Run("2", func(t *testing.T) {
		t.Parallel()
		assert(t, TT{"Sphinx of black quartz judge my vow", "Sphinx of black quartz, judge my vow"})
	})
}

// Diff compares two items and returns a human-readable diff string. If the
// items are equal, the string is empty.
func Diff[T any](got, want T) string {
	opts := cmp.Options{
		cmp.Exporter(func(reflect.Type) bool { return true }),
		cmpopts.EquateEmpty(),
	}
	diff := cmp.Diff(got, want, opts...)
	if diff != "" {
		return "\n-got +want\n" + diff
	}
	return ""
}

// Callers prints the stack trace of everything up til the line where Callers()
// was invoked.
func Callers() string {
	var pc [50]uintptr
	n := runtime.Callers(2, pc[:]) // skip runtime.Callers + Callers
	callsites := make([]string, 0, n)
	frames := runtime.CallersFrames(pc[:n])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callsites = append(callsites, frame.File+":"+strconv.Itoa(frame.Line))
	}
	callsites = callsites[:len(callsites)-1] // skip testing.tRunner
	if len(callsites) == 1 {
		return ""
	}
	var b strings.Builder
	for i := len(callsites) - 1; i >= 0; i-- {
		if b.Len() > 0 {
			b.WriteString(" -> ")
		}
		b.WriteString(filepath.Base(callsites[i]))
	}
	return "\n" + b.String() + ":"
}
