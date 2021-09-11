package main

import "net/http"

type myhandler struct{}

func (m *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {

	//定义handler对象
	mh := myhandler{}

	//定义server变量
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello"))
	//})

	//监听地址
	//http.ListenAndServe("localhost:8080",nil) //nil默认调用DefalutServerMux
}
