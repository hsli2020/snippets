package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main() {
	req := httptest.NewRequest("POST", "/", strings.NewReader(`
		{"email":"name@example.com", "name":"John Doe"}
	`))
	w := httptest.NewRecorder()

	handler := createHandler(UpdateUser)
	handler(w, req)

	io.Copy(os.Stdout, w.Result().Body)
}

// -------------------------------------------
// The RPC handler
// -------------------------------------------

type UpdateUserRequest struct {
	Email string  `json:"email"`
	Name  *string `json:"name"`
}

type UpdateUserResponse struct {
	Status string `json:"status"`
}

func UpdateUser(req UpdateUserRequest) (*UpdateUserResponse, error) {
	// do something with the request

	return &UpdateUserResponse{"ok"}, nil
}

// -------------------------------------------
// The "framework"
// -------------------------------------------

type handlerFunc[Req any, Resp any] func(Req) (Resp, error)

func createHandler[Req any, Resp any](f handlerFunc[Req, Resp]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := f(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(resp)
	}
}
