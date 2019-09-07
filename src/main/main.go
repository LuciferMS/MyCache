package main

import "http"

//缓存服务启动
func main(){
	http.InitServer(3000)
}
