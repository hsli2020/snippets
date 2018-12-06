Extract an uploaded zip file
============================

Extract a HTTP POSTed zip file. req.Body can’t be passed directly to zip.NewReader() as
it doesn’t implement ReadAt(). Wrap it with bytes.NewReader().

body, err := ioutil.ReadAll(req.Body)
if err != nil {
    // err
}

r, err := zip.NewReader(bytes.NewReader(body), req.ContentLength)
if err != nil {
    // err
}
for _, zf := range r.File {
    dst, err := os.Create(zf.Name)
    if err != nil {
        // err
    }
    defer dst.Close()
    src, err := zf.Open()
    if err != nil {
        // err
    }
    defer src.Close()

    io.Copy(dst, src)
}

Get query value in HTTP handler
===============================

req.URL.Query() returns a map which implements Get() method.

func someHandler(w http.ResponseWriter, r *http.Request) {
    query := req.URL.Query()
    val1 := query.Get("key1")
    val2 := query.Get("key2")

Strings are URL-decoded automatically in URL.Query().

When method is GET, req.FormValue() can be also used to get query. However when method is POST,
you need to use req.URL.Query().


Get local IP addresses
======================

Use net.InterfaceAddrs().

addrs, err := net.InterfaceAddrs()
if err != nil {
    panic(err)
}
for i, addr := range addrs {
    fmt.Printf("%d %v\n", i, addr)
}

If you want to know interface names too, use net.Interfaces() to get a list of interfaces first.

list, err := net.Interfaces()
if err != nil {
    panic(err)
}

for i, iface := range list {
    fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
    addrs, err := iface.Addrs()
    if err != nil {
        panic(err)
    }
    for j, addr := range addrs {
        fmt.Printf(" %d %v\n", j, addr)
    }
}

pkcs#7 padding
==============

Functions to add or remove pkcs#7 padding.

import (
    "bytes"
    "fmt"
)

// Appends padding.
func pkcs7Pad(data []byte, blocklen int) ([]byte, error) {
    if blocklen <= 0 {
        return nil, fmt.Errorf("invalid blocklen %d", blocklen)
    }
    padlen := 1
    for ((len(data) + padlen) % blocklen) != 0 {
        padlen = padlen + 1
    }

    pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
    return append(data, pad...), nil
}

// Returns slice of the original data without padding.
func pkcs7Unpad(data []byte, blocklen int) ([]byte, error) {
    if blocklen <= 0 {
        return nil, fmt.Errorf("invalid blocklen %d", blocklen)
    }
    if len(data)%blocklen != 0 || len(data) == 0 {
        return nil, fmt.Errorf("invalid data len %d", len(data))
    }
    padlen := int(data[len(data)-1])
    if padlen > blocklen || padlen == 0 {
        return nil, fmt.Errorf("invalid padding")
    }
    // check padding
    pad := data[len(data)-padlen:]
    for i := 0; i < padlen; i++ {
        if pad[i] != byte(padlen) {
            return nil, fmt.Errorf("invalid padding")
        }
    }

    return data[:len(data)-padlen], nil
}

Read from stdin or file
=======================

Sometimes, you want to specify a file to read with command line argument. Sometimes, you want to
read from stdin. Following is a simple way to do it.

func openStdinOrFile() io.Reader {
    var err error
    r := os.Stdin
    if len(os.Args) > 1 {
        r, err = os.Open(os.Args[1])
        if err != nil {
            panic(err)
        }
    }
    return r
}

func main() {
    r := openStdinOrFile()
    readSomething(r)
}


HTTP Client with Basic Authentication
=====================================

Add HTTP header with SetBasicAuth().

req, err := http.NewRequest("GET", url, nil)
req.SetBasicAuth(user, pass)
cli := &http.Client{}
resp, err := cli.Do(req)

There’s no API for digest authentication in standard library.


Always pass nil to hash.Sum(b []byte)
=====================================

Do you know the difference between following two functions?

func showMd1(b []byte) {
    hash := md5.New()
    md := hash.Sum(b)
    fmt.Printf("%s\n", hex.EncodeToString(md))
}

func showMd2(b []byte) {
    md := md5.Sum(b)
    fmt.Printf("%s\n", hex.EncodeToString(md[:]))
}

The former appends md5 hash to b. Below is what they prints.

showMd1([]byte{ 1, 2, 3, 4, 5, 6 })
-> 010203040506d41d8cd98f00b204e9800998ecf8427e
Hash of empty byte array is appended to 1,2,3,4,5,6.

showMd2([]byte{ 1, 2, 3, 4, 5, 6 })
-> 6ac1e56bc78f031059be7be854522c4c
Hash of []byte{1,2,3,4,5,6}

It’s pretty confusing. hash.Sum() is not the same as final() or digest() in other platforms.
In most cases, you should not pass a slice to hash.Sum(). Pass nil to hash.Sum() just like,

    hash := md5.New()
    hash.Write(b)
    md := hash.Sum(nil)


Create a temporary file
=======================

Use os.TempDir() to get the name of the directory and ioutil.TempFile() to create a file.

file, err := ioutil.TempFile(os.TempDir(), "prefix")
defer os.Remove(file.Name())

It’s caller’s responsibility to remove the file.


Open a file for writing
=======================

Use os.OpenFile()

f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
if err != nil {
    panic(err)
}

If you want to open a file just for reading, os.Open() is easier to use.

f, err := os.Open(filename)
if err != nil {
    panic(err)
}


Read one line from io.Reader
============================

bio.ReadBytes() is more convenient than bio.ReadLine().

bio := bufio.NewReader(os.Stdin)
for {
    line, err := bio.ReadBytes('\n')
    if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }
    sline := strings.TrimRight(string(line), "\n")
    ...
}


Parse HTML
==========

go.net/html is your friend. Create a tokenizer with html.NewTokenizer(io.Reader).
Call Next() to proceed to next token.

import (
    "code.google.com/p/go.net/html"
)

func parseHtml(r io.Reader) {
    d := html.NewTokenizer(r)
    for {
        // token type
        tokenType := d.Next()
        if tokenType == html.ErrorToken {
            return
        }
        token := d.Token()
        switch tokenType {
            case html.StartTagToken: // <tag>
                // type Token struct {
                //     Type     TokenType
                //     DataAtom atom.Atom
                //     Data     string
                //     Attr     []Attribute
                // }
                //
                // type Attribute struct {
                //     Namespace, Key, Val string
                // }
            case html.TextToken: // text between start and end tag
            case html.EndTagToken: // </tag>
            case html.SelfClosingTagToken: // <tag/>

        }
    }
}


Check if file exists
====================

Create FileInfo first and call IsDir().

finfo, err := os.Stat("filename.txt")
if err != nil {
    // no such file or dir
    return
}
if finfo.IsDir() {
    // it's a file
} else {
    // it's a directory
}


Wait until all the background goroutine finish
==============================================

Something like Java’s CountDownLatch

var w sync.WaitGroup
w.Add(2)
go func() {
    // do something
    w.Done()
}
go func() {
    // do something
    w.Done()
}
w.Wait()


AES encryption in CBC mode
==========================

Create cipher.Block with NewCipher() and wrap it using NewCBCEncrypter().

iv := {1,2,3,4,5,... }
block, err := aes.NewCipher(key)
aes := cipher.NewCBCEncrypter(block, iv)
aes.CryptBlocks(out, in)


Get random value
================

Get random value in 1..99

max := big.NewInt(100)
i, err := rand.Int(rand.Reader, max)


HTTP handler func
=================

Example handler func.

func someHandler(w http.ResponseWriter, r *http.Request) {
    // read form value
    value := r.FormValue("value")
    if r.Method == "POST" {
        // receive posted data
        body, err := ioutil.ReadAll(r.Body)
    }
}

func main() {
    http.HandleFunc("/", someHandler)
    http.ListenAndServe(":8080", nil)
}


SHA256
======

Use sha256. If you want to convert to hex string, use hex.EncodeTostring().

hash := sha256.New()
hash.Write(data)
md := hash.Sum(nil)
mdStr := hex.EncodeToString(md)


Add custom HTTP header
======================

Set a header to a request first. Pass the request to a client.

client := &http.Client{]
req, err := http.NewRequest("POST", "http://example.com", bytes.NewReader(postData))
req.Header.Add("User-Agent", "myClient")
resp, err := client.Do(req)
defer resp.Body.Close()


Simple HTTP client
==================

Just call http.Get().

resp, err := http.Get("http://www.google.co.jp")
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)


Base64 encoding
===============

Note: choose the write encoding.

var b bytes.Buffer
w := base64.NewEncoder(base64.URLEncoding, &b)
w.Write(data)
w.Close()
data := b.Bytes()


RC4 encryption
==============

Use rc4 package.

key := []byte{ 1, 2, 3, 4, 5, 6, 7 }
c, err := rc4.NewCipher(key)
dst := make([]byte, len(src))
c.XORKeyStream(dst, src)


Connect two processes with a pipe
=================================

Set a pipe to Cmd.Stdout.

generator := exec.Command("cmd1")
consumer := exec.Command("cmd2")

pipe, err := consumer.StdinPipe()
generator.Stdout = pipe

Read stdout of subprocess
=========================

Create pipe with StdoutPipe()

cmd := exec.Command("cmd", "args")
stdout, err := cmd.StdoutPipe()
cmd.Start()
r := bufio.NewReader(stdout)
line, _, err := r.ReadLine()


Set read timeout to socket
==========================

Set timeout in absolute time. Following code set timeout in 10 seconds.

timeoutSec := 10
sock.SetReadDeadline(time.Now().Add(timeoutSec * time.Second)


Write UDP server
================

Listen with net.ListenUDP()

var buf [1024]byte
addr, err := net.ResolveUDPAddr("udp", ":69")
sock, err := net.ListenUDP("udp", addr)
for {
    rlen, remote, err := sock.ReadFromUDP(buf[:])


Encode binary
=============

Create bytes.Buffer first and write to it with binary package.

buf := new(bytes.Buffer)
binary.Write(buf, binary.BigEndian, value)
binary.Write(buf, binary.BigEndian, value2)
data := buf.Bytes()


Write a UDP client
==================

Use net package

serverAddr, err := net.ResolveUDPAddr("udp", "192.168.1.1:69")
con, err := net.DialUDP("udp", nil, serverAddr)


Write a TCP client
==================

Use net package

serverAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.1:5000")
con, err := net.DialTCP("tcp", nil, serverAddr);


Read a line from stdin
======================

Use bufio and os.Stdin.

bio := bufio.NewReader(os.Stdin)
line, hasMoreInLine, err := bio.ReadLine()

hasMoreInLine will be true if the line is too long for the buffer. It will be false
when returning the last fragment of the line.

