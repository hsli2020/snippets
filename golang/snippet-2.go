
	tpl, err := template.ParseFiles("tpl.gohtml")
	err = tpl.Execute(os.Stdout, nil)
//--

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

//--
	http.Redirect(w, req, "/", http.StatusSeeOther)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
//--
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
//--
	c, err := req.Cookie("my-cookie")

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path: "/",
	})
//--
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.Error(w, http.StatusText(400), http.StatusBadRequest)
//--
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
//--
	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)
//--
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
//--
