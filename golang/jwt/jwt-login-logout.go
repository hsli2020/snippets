package main

import (
 	"fmt"
 	jwt "github.com/dgrijalva/jwt-go"
 	"golang.org/x/crypto/bcrypt"
 	"html/template"
 	"log"
 	"net/http"
 	"time"
)

// used to encrypt/decrypt JWT tokens. Change it to yours.
var jwtTokenSecret = "abc123456def"

const dashBoardPage = `<html><body>

  {{if .Username}}
  <p><b>{{.Username}}</b>, welcome to your dashboard! <a href="/logout">Logout!</a></p>
  {{else}}
  <p>Either your JSON Web token has expired or you've logged out! <a href="/login">Login</a></p>
  {{end}}

  </body></html>`

const logUserPage = `<html><body>
  {{if .LoginError}}<p style="color:red">Either username or password is not in our record! Sign Up?</p>{{end}}

  <form method="post" action="/login">
  {{if .Username}}
  <p><b>{{.Username}}</b>, you're already logged in! <a href="/logout">Logout!</a></p>
  {{else}}
  <label>Username:</label>
  <input type="text" name="Username"><br>

  <label>Password:</label>
  <input type="password" name="Password">

  <span style="font-style:italic"> Enter: mynakedpassword</span><br>
  <input type="submit" name="Login" value="Let me in!">
  {{end}}
  </form>
  </body></html>`

func DashBoardPageHandler(w http.ResponseWriter, r *http.Request) {
 	conditionsMap := map[string]interface{}{}

 	// THIS PAGE should ONLY obe accessible to those users that logged in

 	// check if user already logged in
 	username, _ := ExtractTokenUsername(r)

 	if username != "" {
 		conditionsMap["Username"] = username
 	}

 	if err := dashboardTemplate.Execute(w, conditionsMap); err != nil {
 		log.Println(err)
 	}
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

 	conditionsMap := map[string]interface{}{}

 	// check if user already logged in
 	username, _ := ExtractTokenUsername(r)

 	if username != "" { // user already logged in!
 		conditionsMap["Username"] = username
 		conditionsMap["LoginError"] = false
 	}

 	// verify username and password
 	if r.FormValue("Login") != "" && r.FormValue("Username") != "" {
 		username := r.FormValue("Username")
 		password := r.FormValue("Password")

 		// NOTE: here is where you want to query your database to retrieve the hashed password
 		// for username.
 		// For this tutorial and simplicity sake, we will simulate the retrieved hashed password
 		// as $2a$10$4Yhs5bfGgp4vz7j6ScujKuhpRTA4l4OWg7oSukRbyRN7dc.C1pamu
 		// the plain password is 'mynakedpassword'
 		// see https://www.socketloop.com/tutorials/golang-bcrypting-password for more details
 		// on how to generate bcrypted password

 		hashedPasswordFromDatabase := []byte("$2a$10$4Yhs5bfGgp4vz7j6ScujKuhpRTA4l4OWg7oSukRbyRN7dc.C1pamu")

 		if err := bcrypt.CompareHashAndPassword(hashedPasswordFromDatabase, []byte(password)); err != nil {
 			log.Println("Either username or password is wrong")
 			conditionsMap["LoginError"] = true
 		} else {
 			log.Println("Logged in :", username)
 			conditionsMap["Username"] = username
 			conditionsMap["LoginError"] = false

 			// create a new JSON Web Token and redirect to dashboard
 			tokenString, err := createToken(username)

 			if err != nil {
 				log.Println(err) // of course, this is too simple, your program should prevent login if token cannot be generated!!
 				os.Exit(1)
 			}

 			// create the cookie for client(browser)
 			expirationTime := time.Now().Add(1 * time.Hour) // cookie expired after 1 hour

 			cookie := &http.Cookie{
 				Name:    "token",
 				Value:   tokenString,
 				Expires: expirationTime,
 			}

 			http.SetCookie(w, cookie)

 			// After the cookie is created, the client(browser) will send in the cookie
 			// for every request. Our server side program will unpack the tokenString inside the cookie's Value
 			// for authentication before serving...
 			// This is one big advantage of JWT over session. The burden has been shifted to client instead of taking memory space
 			// on the server side. This helps a lot with the scaling process.

 			http.Redirect(w, r, "/dashboard", http.StatusFound)
 		}

 	}

 	if err := logUserTemplate.Execute(w, conditionsMap); err != nil {
 		log.Println(err)
 	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

 	// For JWT, log out is easy. Just destroy the cookie

 	// see https://golang.org/pkg/net/http/#Cookie
 	// Setting MaxAge<0 means delete cookie now.

 	c := http.Cookie{
 		Name:   "token",
 		MaxAge: -1}
 	http.SetCookie(w, &c)

 	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}

func createToken(username string) (string, error) {
 	claims := jwt.MapClaims{}
 	claims["authorized"] = true
 	claims["username"] = username                            //embed username inside the token string
 	claims["expired"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 	return token.SignedString([]byte(jwtTokenSecret))
}

func ExtractTokenUsername(r *http.Request) (string, error) {

 	// get our token string from Cookie
 	biscuit, err := r.Cookie("token")

 	var tokenString string
 	if err != nil {
 		tokenString = ""
 	} else {
 		tokenString = biscuit.Value
 	}

 	// abort
 	if tokenString == "" {
 		return "", nil
 	}

 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
 		}
 		return []byte(jwtTokenSecret), nil
 	})
 	if err != nil {
 		return "", err
 	}
 	claims, ok := token.Claims.(jwt.MapClaims)
 	if ok && token.Valid {
 		username := fmt.Sprintf("%s", claims["username"]) // convert to string
 		if err != nil {
 			return "", err
 		}
 		return username, nil
 	}
 	return "", nil
}

var dashboardTemplate = template.Must(template.New("").Parse(dashBoardPage))
var logUserTemplate = template.Must(template.New("").Parse(logUserPage))

func main() {
 	fmt.Println("Server starting, point your browser to localhost:8080/login to start")
 	http.HandleFunc("/login", LoginPageHandler)
 	http.HandleFunc("/dashboard", DashBoardPageHandler)
 	http.HandleFunc("/logout", LogoutHandler)
 	http.ListenAndServe(":8080", nil)
}
