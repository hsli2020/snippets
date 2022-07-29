package main // https://github.com/timpratim/urlshortener

//curl -X POST -d "url=https://www.youtube.com/watch?v=4KfuQwB5rIs&t=1s" \
//	http://localhost:8080/shorten

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ShortURL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null"`
	Shortened string `gorm:"not null"`
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ShortURL{})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := ShortenURL(original)
		fmt.Println(shortened)

		// Create a new record in the database.
		db.Create(&ShortURL{Original: original, Shortened: shortened})

		// return something like {"short_url":"http://localhost:8080/abc123"}
		fmt.Fprintf(w, `{"short_url":"%s"}`, shortened)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RedirectURL(db, w, r)
	})

	http.ListenAndServe(":8080", nil)
}

func RedirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	shortened := "http://localhost:8080/" + id

	var url ShortURL
	db.First(&url, "shortened = ?", shortened)

	http.Redirect(w, r, url.Original, http.StatusFound)
}

// create a random string of size 6 from the alphabet
func ShortenURL(url string) string {
	s := ""

	// rand.Intn(26) returns a random number between 0 and 25.
	// 97 is the ascii value of 'a'.
	// So rand.Intn(26) + 97 returns a random lowercase letter.
	for i := 0; i < 6; i++ {
		s += string(rand.Intn(26) + 97)
	}

	shortendURL := fmt.Sprintf("http://localhost:8080/%s", s)
	return shortendURL
}

func ConnectDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	return db, nil
}
