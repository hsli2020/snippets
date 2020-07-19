package main

import (
	"fmt"
	//"encoding/gob"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const SESSID = "GOSESSION"
var store *sessions.FilesystemStore

var tpl *template.Template

func init() {
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	authKey := []byte("super-secret-key")
	encryptionKey := []byte("super-secret-key-1234567")

	store = sessions.NewFilesystemStore("sessions/", authKey, encryptionKey)
	store.MaxLength(32*1024)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	//gob.Register(User{})
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/forbidden", forbidden)
	router.HandleFunc("/secret", secret)
	http.ListenAndServe(":8080", router)
}

// index serves the index html file
func index(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSID)
	checkError(err)

	user := getUser(session)
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

// login authenticates the user
func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSID)
	checkError(err)

	if r.FormValue("code") != "code" {
		if r.FormValue("code") == "" {
			session.AddFlash("Must enter a code")
		}
		session.AddFlash("The code was incorrect")
		err = session.Save(r, w)
		checkError(err)

		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	}

	username := r.FormValue("username")
	session.Values["user"] = username

	err = session.Save(r, w)
	checkError(err)

	http.Redirect(w, r, "/secret", http.StatusFound)
}

// logout revokes authentication for a user
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSID)
	checkError(err)

	session.Values["user"] = ""
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	checkError(err)
	http.Redirect(w, r, "/", http.StatusFound)
}

// secret displays the secret message for authorized users
func secret(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSID)
	checkError(err)

	var list []string

	val := session.Values["list"]
	if val != nil {
		list = val.([]string)
	}

	s := fmt.Sprintf("%x", securecookie.GenerateRandomKey(128))
	list = append(list, s)

	session.Values["list"] = list
	err = session.Save(r, w)
	checkError(err)

	user := getUser(session)

	if user == "" {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		checkError(err)
		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	}

	tpl.ExecuteTemplate(w, "secret.gohtml", user)
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSID)
	checkError(err)

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	checkError(err)

	tpl.ExecuteTemplate(w, "forbidden.gohtml", flashMessages)
}

func getUser(s *sessions.Session) string {
	var user string
	val := s.Values["user"]
	user, ok := val.(string)
	if !ok {
		return ""
	}
	return user
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
