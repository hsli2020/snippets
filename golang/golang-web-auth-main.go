package main  // https://github.com/blessedmadukoma/golang-web-auth/blob/2-backend/main.go

import (
	"database/sql" //new
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os" //new
	"strings"
	"time" //new

	_ "github.com/go-sql-driver/mysql" // new
	"github.com/gorilla/context"
	"github.com/joho/godotenv"   // new
	"golang.org/x/crypto/bcrypt" //new
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

/*
// schema.sql

CREATE DATABASE golangwebauth;

USE DATABASE golangwebauth;

CREATE TABLE user(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  firstname VARCHAR(255) NOT NULL,
  lastname VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  createdDate TIMESTAMP
);
*/

type User struct {
	ID          int
	FirstName   string    `json:"firstname" validate:"required, gte=3"`
	LastName    string    `json:"lastname" validate:"required, gte=3"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"createdDate"`
}

// New
func dbConn() (db *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println(dbDriver, dbUser, dbPass, dbName)

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB Connected!!")
	return db
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := dbConn()
		firstName := r.FormValue("FirstName")
		lastName := r.FormValue("LastName")
		email := r.FormValue("email")
		fmt.Printf("%s, %s, %s\n", firstName, lastName, email)

		password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			tpl.ExecuteTemplate(w, "Register", err)
		}

		dt := time.Now()

		createdDateString := dt.Format("2006-01-02 15:04:05")

		// Convert the time before inserting into the database
		createdDate, err := time.Parse("2006-01-02 15:04:05", createdDateString)
		if err != nil {
			log.Fatal("Error converting the time:", err)
		}

		_, err = db.Exec("
			INSERT INTO user (firstname, lastname,email,password,createdDate)
			     VALUES (?,?,?,?,?)",
			     firstName, lastName, email, password, createdDate)
		if err != nil {
			fmt.Println("Error when inserting: ", err.Error())
			panic(err.Error())
		}
		log.Println("=> Inserted: First Name: " + firstName + " | Last Name: " + lastName)

		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	} else if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "register.html", nil)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := dbConn()
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Printf("%s, %s\n", email, password)

		if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
			fmt.Println("Parameter's can't be empty")
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}

		checkUser, err := db.Query("
			SELECT id, createdDate, password, firstname, lastname, email
			  FROM user 
			 WHERE email=?", email)
		if err != nil {
			panic(err.Error())
		}
		user := &User{}
		for checkUser.Next() {
			var id int
			var password, firstName, lastName, email string
			var createdDate time.Time
			err = checkUser.Scan(&id, &createdDate, &password, &firstName, &lastName, &email)
			if err != nil {
				panic(err.Error())
			}
			user.ID = id
			user.FirstName = firstName
			user.LastName = lastName
			user.Email = email
			user.Password = password
			user.CreatedDate = createdDate
		}

		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			fmt.Println(errf)
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		} else {
			tpl.ExecuteTemplate(w, "dashboard.html", user)
			return
		}
	} else if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "login.html", nil)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dashboard.html", nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logouth", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on: http://localhost:8000")
	err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)) // context to prevent memory leak
	if err != nil {
		log.Fatal(err)
	}
}
