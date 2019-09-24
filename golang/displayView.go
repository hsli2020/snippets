package main

import (
	"IMtest/app/controller"
	"html/template"
	"log"
	"net/http"
)

//万能渲染模版
func displayView() {
	tpl, err := template.ParseGlob("./app/view/**/*")
	if err != nil {
		//打印并直接退出
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
}

func main() {
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/addfriend", controller.AddFriend)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/createcommunity", controller.CreateCommunity)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/chat", controller.Chat)
	http.HandleFunc("/attach/upload", controller.FileUpload)

	//提供静态资源目录支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/resource/", http.FileServer(http.Dir(".")))

	displayView()
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
