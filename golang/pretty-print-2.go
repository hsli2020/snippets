package pretty

// https://github.com/domonda/go-pretty

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"unicode/utf8"
)

var (
	typeOfByte = reflect.TypeOf(byte(0))

// typeOfSortInterface = reflect.TypeOf((*sort.Interface)(nil)).Elem()
)

// Println pretty prints a value to os.Stderr followed by a newline
func Println(value interface{}) {
	Print(value)
	os.Stderr.Write([]byte{'\n'})
}

// Print pretty prints a value to os.Stderr
func Print(value interface{}) {
	fprintValue(os.Stderr, value)
}

// Fprint pretty prints a value to a io.Writer
func Fprint(w io.Writer, value interface{}) {
	fprintValue(w, value)
}

// Fprint pretty prints a value to a io.Writer followed by a newline
func Fprintln(w io.Writer, value interface{}) {
	Fprint(w, value)
	os.Stderr.Write([]byte{'\n'})
}

// Sprint pretty prints a value to a string
func Sprint(value interface{}) string {
	var b strings.Builder
	fprintValue(&b, value)
	return b.String()
}

func fprintValue(w io.Writer, value interface{}) {
	if value == nil {
		fmt.Fprint(w, "nil")
	} else {
		fprint(w, reflect.ValueOf(value))
	}
}

func fprint(w io.Writer, v reflect.Value) {
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	t := v.Type()

	switch t.Kind() {
	case reflect.Ptr:
		// Pointers were dereferenced above, so only nil left as possibility
		fmt.Fprint(w, "nil")

	case reflect.String:
		fmt.Fprintf(w, "%q", v.Interface())

	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		fmt.Fprint(w, v.Interface())

	case reflect.Float32, reflect.Float64:
		fmt.Fprint(w, v.Interface())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprint(w, v.Interface())

	case reflect.Array:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{','})
			}
			fprint(w, v.Index(i))
		}
		w.Write([]byte{']'})

	case reflect.Slice:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		if t.Elem() == typeOfByte && utf8.Valid(v.Bytes()) {
			fmt.Fprintf(w, "%q", v.Interface())
			return
		}
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{','})
			}
			fprint(w, v.Index(i))
		}
		w.Write([]byte{']'})

	case reflect.Map:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		// TODO sort map if possible
		// if t.Key().Implements(typeOfSortInterface) {
		// 	// TODO Need to make a temp sorted copy
		// }
		// switch t.Key().Kind() {
		// case reflect.String:
		// case reflect.Slice, reflect.Array:
		// }
		fmt.Fprintf(w, "%s{", t.Name())
		for i, iter := 0, v.MapRange(); iter.Next(); i++ {
			if i > 0 {
				w.Write([]byte{','})
			}
			fprint(w, iter.Key())
			w.Write([]byte{':'})
			fprint(w, iter.Value())
		}
		w.Write([]byte{'}'})

	case reflect.Struct:
		fmt.Fprintf(w, "%s{", t.Name())
		for i := 0; i < t.NumField(); i++ {
			if i > 0 {
				w.Write([]byte{','})
			}
			f := t.Field(i)
			if !f.Anonymous {
				fmt.Fprintf(w, "%s:", f.Name)
			}
			fprint(w, v.Field(i))
		}
		w.Write([]byte{'}'})

	case reflect.Chan:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		switch t.ChanDir() {
		case reflect.RecvDir:
			fmt.Fprint(w, "<-chan ", t.Elem().String())
		case reflect.SendDir:
			fmt.Fprint(w, "chan<- ", t.Elem().String())
		case reflect.BothDir:
			fmt.Fprint(w, "chan ", t.Elem().String())
		}

	case reflect.Func:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		fmt.Fprint(w, "func")

	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		fmt.Fprint(w, "interface{}")

	case reflect.UnsafePointer:
		fmt.Fprint(w, v.Interface())

	default:
		panic("invalid kind")
	}
}
