// go run json.go actions.go user.go httpauth.go ful.go request.go

// ful.go /** main program, does the wiring up */
package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/ful")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := connectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/users/:id", Authenticate(db, UserShow(db)))
	router.GET("/search/users", Authenticate(db, UserSearch(db)))
	router.POST("/users", Authenticate(db, UserCreate(db)))
	router.PUT("/users/:id", Authenticate(db, UserUpdate(db)))
	router.DELETE("/users/:id", Authenticate(db, UserRemove(db)))

	log.Print("Starting....\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// json.go /** JSON-related helpers */
package main

import (
	"encoding/json"
	"net/http"
)

type JsonError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func MarkAsJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SendJSON(w http.ResponseWriter, o interface{}) error {
	if json, err := json.Marshal(o); err == nil {
		w.Write(json)
		return nil
	} else {
		return err
	}
}

func SendError(w http.ResponseWriter, err error) {
	w.WriteHeader(400)                       // invalid request
	SendJSON(w, JsonError{400, err.Error()}) // XXX check error here?
}

// actions.go /** http user actions */
package main

import (
	"net/http"
	"database/sql"

	"github.com/julienschmidt/httprouter"
)

type authHandler func(*User, http.ResponseWriter, *http.Request, httprouter.Params)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func UserShow(db *sql.DB) authHandler {
	return func(user *User, w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
		MarkAsJSON(w)
		if user, err := GetUserBy(db, "id", p.ByName("id")); err != nil {
			w.WriteHeader(404)
		} else if user.Role == "admin" {
			w.WriteHeader(401)
			SendJSON(w, JsonError{401, "unauthorized"})
		} else {
			SendJSON(w, user)
		}
	}
}

func UserSearch(db *sql.DB) authHandler {
	return func(u *User, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		MarkAsJSON(w)
		email := r.URL.Query().Get("q")
		if len(email) == 0 {
			w.WriteHeader(404)
		} else if user, err := GetUserBy(db, "email", email); err != nil {
			w.WriteHeader(404)
		} else {
			SendJSON(w, []User{user})
		}
	}
}

func UserCreate(db *sql.DB) authHandler {
	return func(u *User, w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		MarkAsJSON(w)
		user, err := parseUser(r)
		if err != nil {
			SendError(w, err)
			return
		}
		user.Role = validateUserRole(user.Role, u)
		newUser, err := InsertUser(db, user)
		if err != nil {
			SendError(w, err)
			return
		}
		MarkAsJSON(w)
		w.WriteHeader(201)
		SendJSON(w, newUser)
	}
}

func UserUpdate(db *sql.DB) authHandler {
	return func(u *User, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		MarkAsJSON(w)
		id := p.ByName("id")
		user, err := parseUser(r)
		if err != nil {
			SendError(w, err)
			return
		}
		if updateUser, err := GetUserBy(db, "id", id); err != nil || updateUser.Id == 0 {
			// no such user
			w.WriteHeader(404)
			return
		}
		if user.Role == "normal" || user.Role == "admin" {
			if u.Role == "admin" {
				UpdateUserColumn(db, id, "role", user.Role)
			} else {
				// non-admin can't update other users' roles
				w.WriteHeader(401)
				return
			}
		}
		if user.Lastname != "" {
			UpdateUserColumn(db, id, "lastname", user.Lastname)
		}
		if user.Firstname != "" {
			UpdateUserColumn(db, id, "firstname", user.Firstname)
		}
		if user.Email != "" {
			UpdateUserColumn(db, id, "email", user.Email)
		}
		if user.Password != "" {
			UpdateUserColumn(db, id, "password", user.Password)
		}
		w.WriteHeader(204)
	}
}

func UserRemove(db *sql.DB) authHandler {
	return func(u *User, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := DeleteUser(db, p.ByName("id")); err != nil {
			SendError(w, err)
			return
		}
	}
}

// user.go /** DB-related query functions */
package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id        int    `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role"`
}

type NewUser struct {
	Id int64 `json:"id"`
}

func GetUserBy(db *sql.DB, column string, value string) (User, error) {
	var user User
	sql := "SELECT id, firstname, lastname, email, password, role FROM user WHERE " + 
		column + " = ?"
	err := db.QueryRow(sql, value).Scan(&user.Id, &user.Firstname, &user.Lastname, 
		&user.Email, &user.Password, &user.Role)
	return user, err
}

func InsertUser(db *sql.DB, user User) (*NewUser, error) {
	sql := `INSERT INTO user (firstname, lastname, email, password, role) VALUES (?, ?, ?, ?, ?)`
	res, err := db.Exec(sql, user.Firstname, user.Lastname, user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	newUser := NewUser{id}
	return &newUser, nil
}

func UpdateUserColumn(db *sql.DB, id, column, value string) error {
	sql := fmt.Sprintf(`UPDATE user SET %s = ? WHERE id = ?`, column)
	_, err := db.Exec(sql, value, id)
	return err
}

func DeleteUser(db *sql.DB, id string) error {
	sql := "DELETE FROM user WHERE id = ?"
	_, err := db.Exec(sql, id)
	return err
}

func validateUserRole(role string, authUser *User) string {
    if authUser.Role != "admin" {
        return "normal"
    }
    if role == "normal" || role == "admin" {
        return role
    }
    return "normal"
}

// httpauth.go /** http-related helpers */
package main

import (
	"encoding/base64"
    "strings"
)

type httpCode int

type BasicUser struct {
    Username, Password string
}

func parseAuthHeader(authHeader []string) (*BasicUser, httpCode) {
    // validate we have the header... (not present: unauthorized)
    if len(authHeader) == 0 {
        return nil, 401
    }
    // validate the header's shape... (bad shape: bad request)
    auth := strings.SplitN(authHeader[0], " ", 2)
    if len(auth) != 2 || auth[0] != "Basic" {
        return nil, 400
    }

    // parse the header... validate the size (bad shape: bad request)
    payload, _ := base64.StdEncoding.DecodeString(auth[1])
    pair := strings.SplitN(string(payload), ":", 2)
    if len(pair) != 2 {
        return nil, 400
    }
    return &BasicUser{pair[0], pair[1]}, 200
}

// request.go /** request-parsing helpers */
package main

import (
	"encoding/json"
	"database/sql"
	"io/ioutil"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Authenticate(db *sql.DB, action authHandler) 
		func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		basicUser, errCode := parseAuthHeader(r.Header["Authorization"])
		if errCode != 200 {
		    w.WriteHeader(int(errCode))
		    return
		}

		user, err := GetUserBy(db, "email", basicUser.Username)
		if err != nil || user.Password != basicUser.Password {
			w.WriteHeader(401)
			return
		}

		action(&user, w, r, p)
	}
}

func parseUser(r *http.Request) (User, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return User{}, err
	}
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, err
	}
	return user, err
}
