package http

import (
	"Cache"
	"net/http"
	"strconv"
)

type Handler func(w http.ResponseWriter, r *http.Request)

var data = Cache.InitDataMap()

func InitServer(port int) {
	handlerMapper := initHandlerWrapper();
	for name, handler := range handlerMapper {
		http.HandleFunc(name, handler)
	}
	http.ListenAndServe("127.0.0.1:" + strconv.Itoa(port), nil)
}

func initHandlerWrapper()(handlerMapper map[string]Handler){
	mapper := make(map[string]Handler)
	mapper["/set"] = setHandler()
	mapper["/get"] = getHandler()
	mapper["/del"] = delHandler()
	return mapper
}

func setHandler()(handler Handler){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query();
		result := data.Set(vars["key"][0], vars["value"][0])
		data.Print()
		w.Write(result.GetData())
	}
}

func getHandler()(handler Handler){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		result := data.Get(vars["key"][0])
		w.Write(result.GetData())
	}
}

func delHandler()(handler Handler){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		result := data.Del(vars["key"][0])
		data.Print()
		w.Write(result.GetData())
	}
}