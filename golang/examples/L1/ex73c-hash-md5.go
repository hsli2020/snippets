
package main

import "fmt"
import "io"
import "crypto/md5"

func md(str string) string {
	h := md5.New()
	io.WriteString(h, str)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Println(md("Hello, playground"))
}
