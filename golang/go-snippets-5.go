
import "database/sql"

func connectDB() (*sql.DB, error) {
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
	query := "SELECT id, firstname, lastname, email, password, role FROM user WHERE " + column + " = ?"
	err := db.QueryRow(query, value).Scan(
        &user.Id,
        &user.Firstname,
        &user.Lastname,
        &user.Email,
        &user.Password,
        &user.Role)
	return user, err
}

func InsertUser(db *sql.DB, user User) (*NewUser, error) {
	query := `INSERT INTO user (firstname, lastname, email, password, role)
        VALUES (?, ?,             ?,        ?,    ?)`

	res, err := db.Exec(query, user.Firstname, user.Lastname, user.Email, user.Password, user.Role)
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
    query := fmt.Sprintf(`UPDATE user SET %s = ?  WHERE id = ?`, column)
    _, err := db.Exec(query, value, id)
    return err
}

func DeleteUser(db *sql.DB, id string) error {
    query := "DELETE FROM user WHERE id = ?"
    _, err := db.Exec(query, id)
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
