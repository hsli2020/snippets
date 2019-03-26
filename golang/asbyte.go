// Command asbyte prints the contents of files as Go byte slices.
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode"
)

var re = regexp.MustCompile(`\W`)

func main() {
	fmt.Println("// nolint: lll\nvar (")
	for _, f := range os.Args[1:] {
		fp, err := os.Open(f)
		fatal(err)

		// Make sure it's a valid variable name. Don't try to be too smart, as
		// they will probably need to be renamed manually anyway.
		fn := re.ReplaceAllString(f, "_")
		if unicode.IsDigit(rune(fn[0])) {
			fn = "x" + fn
		}

		fmt.Printf("\t%s = []byte{", fn)
		for {
			buf := make([]byte, 1024)
			n, err := fp.Read(buf)
			if err != nil {
				_ = fp.Close()
				if err == io.EOF {
					break
				}
				fatal(err)
			}

			for _, c := range buf[:n] {
				fmt.Printf("%#x, ", c)
			}
		}
		fmt.Println("}")
	}
	fmt.Println(")")
}

func fatal(err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, "asbyte: %v", err)
	os.Exit(1)
}
