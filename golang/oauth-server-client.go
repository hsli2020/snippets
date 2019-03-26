package main

// https://tutorialedge.net/golang/go-oauth2-tutorial/

import (
    "log"
    "net/http"
    "net/url"
    "os"

    "github.com/go-session/session"
    "gopkg.in/oauth2.v3/errors"
    "gopkg.in/oauth2.v3/manage"
    "gopkg.in/oauth2.v3/models"
    "gopkg.in/oauth2.v3/server"
    "gopkg.in/oauth2.v3/store"
)

func main() {
    manager := manage.NewDefaultManager()
    // token store
    manager.MustTokenStorage(store.NewMemoryTokenStore())

    clientStore := store.NewClientStore()
    clientStore.Set("222222", &models.Client{
        ID:     "222222",
        Secret: "22222222",
        Domain: "http://localhost:9094",
    })
    manager.MapClientStorage(clientStore)

    srv := server.NewServer(server.NewConfig(), manager)
    srv.SetUserAuthorizationHandler(userAuthorizeHandler)

    srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
        log.Println("Internal Error:", err.Error())
        return
    })

    srv.SetResponseErrorHandler(func(re *errors.Response) {
        log.Println("Response Error:", re.Error.Error())
    })

    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/auth", authHandler)

    http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
        err := srv.HandleAuthorizeRequest(w, r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
    })

    http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
        err := srv.HandleTokenRequest(w, r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    log.Println("Server is running at 9096 port.")
    log.Fatal(http.ListenAndServe(":9096", nil))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
    store, err := session.Start(nil, w, r)
    if err != nil {
        return
    }

    uid, ok := store.Get("UserID")
    if !ok {
        if r.Form == nil {
            r.ParseForm()
        }
        store.Set("ReturnUri", r.Form)
        store.Save()

        w.Header().Set("Location", "/login")
        w.WriteHeader(http.StatusFound)
        return
    }
    userID = uid.(string)
    store.Delete("UserID")
    store.Save()
    return
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    store, err := session.Start(nil, w, r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if r.Method == "POST" {
        store.Set("LoggedInUserID", "000000")
        store.Save()

        w.Header().Set("Location", "/auth")
        w.WriteHeader(http.StatusFound)
        return
    }
    outputHTML(w, r, "static/login.html")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
    store, err := session.Start(nil, w, r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if _, ok := store.Get("LoggedInUserID"); !ok {
        w.Header().Set("Location", "/login")
        w.WriteHeader(http.StatusFound)
        return
    }

    if r.Method == "POST" {
        var form url.Values
        if v, ok := store.Get("ReturnUri"); ok {
            form = v.(url.Values)
        }
        u := new(url.URL)
        u.Path = "/authorize"
        u.RawQuery = form.Encode()
        w.Header().Set("Location", u.String())
        w.WriteHeader(http.StatusFound)
        store.Delete("Form")

        if v, ok := store.Get("LoggedInUserID"); ok {
            store.Set("UserID", v)
        }
        store.Save()

        return
    }
    outputHTML(w, r, "static/auth.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
    file, err := os.Open(filename)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer file.Close()
    fi, _ := file.Stat()
    http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

// Our Client

package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "golang.org/x/oauth2"
)

var (
    config = oauth2.Config{
        ClientID:     "222222",
        ClientSecret: "22222222",
        Scopes:       []string{"all"},
        RedirectURL:  "http://localhost:9094/oauth2",
        // This points to our Authorization Server
        // if our Client ID and Client Secret are valid
        // it will attempt to authorize our user
        Endpoint: oauth2.Endpoint{
            AuthURL:  "http://localhost:9096/authorize",
            TokenURL: "http://localhost:9096/token",
        },
    }
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Homepage Hit!")
    u := config.AuthCodeURL("xyz")
    http.Redirect(w, r, u, http.StatusFound)
}

// Authorize
func Authorize(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    state := r.Form.Get("state")
    if state != "xyz" {
        http.Error(w, "State invalid", http.StatusBadRequest)
        return
    }

    code := r.Form.Get("code")
    if code == "" {
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

    token, err := config.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    e := json.NewEncoder(w)
    e.SetIndent("", "  ")
    e.Encode(*token)
}

func main() {

    // 1 - We attempt to hit our Homepage route
    // if we attempt to hit this unauthenticated, it
    // will automatically redirect to our Auth
    // server and prompt for login credentials
    http.HandleFunc("/", HomePage)

    // 2 - This displays our state, code and
    // token and expiry time that we get back
    // from our Authorization server
    http.HandleFunc("/oauth2", Authorize)

    // 3 - We start up our Client on port 9094
    log.Println("Client is running at 9094 port.")
    log.Fatal(http.ListenAndServe(":9094", nil))
}