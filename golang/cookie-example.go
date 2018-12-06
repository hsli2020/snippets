package main

import (
	"encoding/json"
	"flag"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "server address")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/get", getCookie)
	mux.HandleFunc("/delete", deleteCookie)
	mux.HandleFunc("/set", setCookie)
	http.ListenAndServe(*addr, mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="#" onclick="alert(document.cookie)">Click here!</a>`))
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("this_is_a_test_cookie")
	if err != nil {
		w.Write([]byte("��ȡcookieʧ��: " + err.Error()))
	} else {
		data, _ := json.MarshalIndent(c, "", "\t")
		w.Write([]byte("��ȡ��cookieֵ: \n" + string(data)))
	}
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "this_is_a_test_cookie",
		MaxAge: -1}
	http.SetCookie(w, &c)
	w.Write([]byte("cookie�ѱ�ɾ��"))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "this_is_a_test_cookie",
		Value:    "true",
		HttpOnly: true,
		//Secure:   true,
		MaxAge: 300}
	http.SetCookie(w, &c)
	w.Write([]byte("cookie�Ѵ���\n"))
}