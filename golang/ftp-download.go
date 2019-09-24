package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

const (
	monthOffset = -1
)

func main() {
	c, err := ftp.Dial(
		os.Getenv("FTP_HOST"),
		ftp.DialWithTimeout(5*time.Second),
		Using ftp.DialWithDebugOutput(os.Stdout)
	)
	if err != nil {
		log.Fatalln("unable to connect: ", err)
	}

	err = c.Login(
		os.Getenv("ACCOUNT"),
		os.Getenv("PASSWORD"),
	)
	if err != nil {
		log.Fatal("login failed: ", err)
	}

	t := time.Now().AddDate(0, monthOffset, 0)
	year, month, _ := t.Date()
	path := fmt.Sprintf("%d/%02d/", year, month)

	entries, err := c.NameList(path)
	if err != nil {
		log.Fatalf("unable to NLIST %s: %v\n", path, err)
	}

	for _, entry := range entries {
		fmt.Printf("retrieving %s...\n", entry)
		resp, err := c.Retr(entry)
		if err != nil {
			log.Fatalf("unable to retrieve %s: %v\n", entry, err)
		}
		defer resp.Close()

		if _, err := os.Stat(entry); os.IsNotExist(err) {
			os.MkdirAll("data/"+path, 0700)
		}
		destination, err := os.Create("data/" + entry)
		if err != nil {
			log.Fatalln("failed to create file:", err)
		}

		b, err := io.Copy(destination, resp)
		if err != nil {
			log.Fatalln("unable to copy response to dest file: ", err)
		}
		fmt.Printf("successfully copied: %s, bytes copied: %d\n", destination.Name(), b)
	}
}

// https://forum.golangbridge.org/t/ftp-entering-extended-passive-mode/14778

// 229 Entering Extended Passive Mode (|||44022|).

// The issue here was I was calling defer resp.Close() within the for loop 
// that was calling c.Retr(entry).
//
// The solution was to call resp.Close() without defer after copy