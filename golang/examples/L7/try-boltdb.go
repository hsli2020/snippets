package main

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/boltdb/bolt"
)

var world = []byte("world")

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    key := []byte("hello")
    value := []byte("Hello World!")

    // store some data
    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(world)
        if err != nil {
            return err
        }

        err = bucket.Put(key, value)
        if err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    // retrieve the data
    err = db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket(world)
        if bucket == nil {
            return fmt.Errorf("Bucket %q not found!", world)
        }

        val := bucket.Get(key)
        fmt.Println(string(val))

        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

	user := &User{"hsli", 48, "Toronto", "pswdhsli", "Don Mills Road"}
	user.save(db)
}

type User struct {
    Name     string
    Age      int
    Location string
    Password string
    Address  string
}

var usersBucket = []byte("users")

func (user *User) save(db *bolt.DB) error {
    // Store the user model in the user bucket using the username as the key.
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists(usersBucket)
        if err != nil {
            return err
        }

        encoded, err := json.Marshal(user)
        if err != nil {
            return err
        }
        return b.Put([]byte(user.Name), encoded)
    })
    return err
}
