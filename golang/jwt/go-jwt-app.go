package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Secret key to uniquely sign the token
var key []byte

// Credential User's login information
type Credential struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token jwt Standard Claim Object
type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create a dummy local db instance as a key value pair
var userdb = map[string]string{
	"user1": "password123",
}

// assign the secret key to key variable on program's first run
func init() {
	// Load the .env file to access the environment variable
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// read the secret_key from the .env file
	key = []byte(os.Getenv("SECRET_KEY"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/me", dashboard).Methods("GET")

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// login user login function
func login(w http.ResponseWriter, r *http.Request) {
	// create a Credentials object
	var creds Credential
	// decode json to struct
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// verify if user exist or not
	userPassword, ok := userdb[creds.Username]

	// if user exist, verify the password
	if !ok || userPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a token object
	var tokenObj = Token {
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// Enter expiration in milisecond
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenObj )

	tokenString, err := token.SignedString(key)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(tokenString)
}

// dashboard User's personalized dashboard
func dashboard(w http.ResponseWriter, r *http.Request) {
	// get the bearer token from the reuest header
	bearerToken := r.Header.Get("Authorization")

	// validate token, it will return Token and error
	token, err := ValidateToken(bearerToken)

	if err != nil {
		// check if Error is Signature Invalid Error
		if err == jwt.ErrSignatureInvalid {
			// return the Unauthorized Status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Return the Bad Request for any other error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Validate the token if it expired or not
	if !token.Valid {
		// return the Unauthoried Status for expired token
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Type cast the Claims to *Token type
	user := token.Claims.(*Token)

	// send the username Dashboard message
	json.NewEncoder(w).Encode(fmt.Sprintf("%s Dashboard", user.Username))
}

// ValidateToken validates the token with the secret key and return the object
func ValidateToken(bearerToken string) (*jwt.Token, error) {
	// format the token string
	tokenString := strings.Split(bearerToken, " ")[1]

	// Parse the token with tokenObj
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token)(interface{}, error) {
		return key, nil
	})

	// return token and err
	return token, err
}