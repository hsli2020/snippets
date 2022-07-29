func Must[T any](x T, err error) T {
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    return x
}

// Before
func main() {
    r, err := regexp.Compile("[")  // error
    if err != nil {
        panic(err)
    }
    fmt.Println(r)
}

// After
func main() {
    r = Must(regexp.Compile("["))  // error
    fmt.Println(r)
}

// Before
func main() {
    src := "./template.txt"
    dst := "./out/template.txt"

    r, err := os.Open(src)
    if err != nil {
        panic(err)
    } 
    defer r.Close()

    w, err := os.Create(dst)
    if err != nil {
        panic(err)
    } 
    defer w.Close()

    if _, err := io.Copy(w, r); err != nil {
        panic(err)
    }

    if err := w.Close(); err != nil {
        panic(err)
    }
}

// After
func main() {
    src := "./template.txt"
    dst := "./out/template.txt"

    r := Must(os.Open(src))
    defer r.Close()

    w := Must(os.Create(dst))
    defer w.Close()

    Must(io.Copy(w, r))

    checkErr(w.Close())
}
