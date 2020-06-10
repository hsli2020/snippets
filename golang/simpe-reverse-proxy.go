// golang-实现反向代理

// 反向代理

// 反向代理在计算机网络中是代理服务器的一种。服务器根据客户端的请求，从其关系的一组或多组后端服务器
// （如Web服务器）上获取资源，然后再将这些资源返回给客户端，客户端只会得知反向代理的IP地址，而不知道在
// 代理服务器后面的服务器集群的存在

// golang最简单实现反向代理

// golang源码httputils中有对于反向代理的实现，最简单的代理甚至可以一行代码实现。

// 1、我们首先开启一个web服务器监听127.0.0.1:8999端口

import (
	"log"
	_ "net/http"
)

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprint(w, "Welcome!")
}

func startServer() {
	err := http.ListenAndServe(":8999", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}

// 2、实现一个简单的代理服务器

// 下面的程序仍然是一个web服务器，监听8888端口，但是其使用了反向代理，因此对:8888的访问都会转发到:8999，
// 输出“Welcome！”。

// 核心的操作在于httputil.NewSingleHostReverseProxy 具有serveHttp方法，此方法对request请求进行了重新
// 封装，并且proxy将得到的response转发给client。

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	host string
	port string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url:%+v\n",r.URL)
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	fmt.Println("hosr::",remote.Host)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{host: "127.0.0.1", port: "8999"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}