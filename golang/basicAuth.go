package main

import (
    "crypto/sha256"
    "crypto/subtle"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

type application struct {
    auth struct {
        username string
        password string
    }
}

func main() {
    app := new(application)

    app.auth.username = os.Getenv("AUTH_USERNAME")
    app.auth.password = os.Getenv("AUTH_PASSWORD")

    if app.auth.username == "" {
        log.Fatal("basic auth username must be provided")
    }

    if app.auth.password == "" {
        log.Fatal("basic auth password must be provided")
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/unprotected", app.unprotectedHandler)
    mux.HandleFunc("/protected", app.basicAuth(app.protectedHandler))

    srv := &http.Server{
        Addr:         ":4000",
        Handler:      mux,
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    log.Printf("starting server on %s", srv.Addr)
    err := srv.ListenAndServeTLS("./localhost.pem", "./localhost-key.pem")
    log.Fatal(err)
}

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the protected handler")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the unprotected handler")
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

// middleware

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract the username and password from the request 
        // Authorization header. If no Authentication header is present 
        // or the header value is invalid, then the 'ok' return value 
        // will be false.
		username, password, ok := r.BasicAuth()
		if ok {
            // Calculate SHA-256 hashes for the provided and expected
            // usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte("your expected username"))
			expectedPasswordHash := sha256.Sum256([]byte("your expected password"))

            // Use the subtle.ConstantTimeCompare() function to check if 
            // the provided username and password hashes equal the  
            // expected username and password hashes. ConstantTimeCompare
            // will return 1 if the values are equal, or 0 otherwise. 
            // Importantly, we should to do the work to evaluate both the 
            // username and password before checking the return values to 
            // avoid leaking information.
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

            // If the username and password are correct, then call
            // the next handler in the chain. Make sure to return 
            // afterwards, so that none of the code below is run.
			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

        // If the Authentication header is not present, is invalid, or the
        // username or password is wrong, then set a WWW-Authenticate 
        // header to inform the client that we expect them to use basic
        // authentication and send a 401 Unauthorized response.
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

// client

package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

func main() {
    client := http.Client{Timeout: 5 * time.Second}

    req, err := http.NewRequest(http.MethodGet, "https://localhost:4000/protected", http.NoBody)
    if err != nil {
        log.Fatal(err)
    }

    req.SetBasicAuth("alice", "p8fnxeqj5a7zbrqp")

    res, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Status: %d\n", res.StatusCode)
    fmt.Printf("Body: %s\n", string(resBody))
}