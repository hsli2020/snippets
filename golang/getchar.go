https://stackoverflow.com/questions/14094190/golang-function-similar-to-getchar

C's getchar() example:

#include <stdio.h>
void main()
{
    char ch;
    ch = getchar();
    printf("Input Char Is :%c",ch);
}

Go equivalent:

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')

    fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))

    // fmt.Printf("You entered: %v", []byte(input))
}


package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    r := bufio.NewReader(os.Stdin)
    c, err := r.ReadByte()
    if err != nil {
        panic(err)
    }
    fmt.Println(c)
}


Assuming that you want unbuffered input (without having to hit enter), this does the job on UNIX systems:

package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
    // restore the echoing state when exiting
    defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

    var b []byte = make([]byte, 1)
    for {
        os.Stdin.Read(b)
        fmt.Println("I got the byte", b, "("+string(b)+")")
    }
}



package main

import (
    "bytes"
    "fmt"

    "github.com/pkg/term"
)

func getch() []byte {
    t, _ := term.Open("/dev/tty")
    term.RawMode(t)
    bytes := make([]byte, 3)
    numRead, err := t.Read(bytes)
    t.Restore()
    t.Close()
    if err != nil {
        return nil
    }
    return bytes[0:numRead]
}

func main() {
    for {
        c := getch()
        switch {
        case bytes.Equal(c, []byte{3}):
            return
        case bytes.Equal(c, []byte{27, 91, 68}): // left
            fmt.Println("LEFT pressed")
        default:
            fmt.Println("Unknown pressed", c)
        }
    }
    return
}


package climenu

import "github.com/pkg/term"

// Returns either an ascii code, or (if input is an arrow) a Javascript key code.
func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}


