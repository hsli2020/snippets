package main

import (
	//"fmt"
	"io"
	"log"
	"os"
)

const logfile = "test.log"

func main() {
	//if !FileExists(logfile) {
	//	CreateFile(logfile)
	//}

	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//log.Fatal(err)
		log.Fatalf("error: %v", err)
	}
	defer f.Close()

	//log.SetOutput(file)
	log.SetOutput(io.MultiWriter(os.Stderr, f))
	log.Println("Hello, logfile")
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
