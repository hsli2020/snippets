package main  // 一种简单的golang web app组织方式 (flat structure)
              // 其原则和目的是，使事情简单化，而不是复杂化
import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	// 所有模版文件存放在templates/目录下，没有子目录
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user/login", userLogin)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Index"
	data["User"] = "Sam Wang"
	tpl.ExecuteTemplate(w, "home-index.gohtml", data)
}

func userLogin(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "User Login"
	tpl.ExecuteTemplate(w, "user-login.gohtml", data)
}

// ##### 模版文件之间采用组合方式

// home-index.html
{{template "html-start.gohtml" .}}
<h1>File: home-index.gohtml</h1>
{{template "html-end.gohtml" .}}

// user-login.gohtml
{{template "html-start.gohtml" .}}
<h1>File: user-login.gohtml</h1>
{{template "html-end.gohtml" .}}

// html-start.gohtml
<html>
<head>
<meta charset="utf-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0"/>
<title>{{.Title}}</title>
<link rel="shortcut icon" type="image/x-icon" href="https://static.oschina.net/new-osc/img/favicon.ico"/>
<link type="text/css" rel="stylesheet" href="/assets/css/semantic.min.css/">
<link type="text/css" rel="stylesheet" href="/assets/css/style.css">
</head>
<body>

// html-end.gohtml
<script src="https://code.jquery.com/jquery-1.11.3.js"></script>
<script>var code = "More...";</script>
</body>
</html>
