package main

import (
	"net/http"
	"log"
	"thinkgo/access"
	_ "thinkgo/controller"
)


func main(){
	//接入http
	http.HandleFunc("/", access.HttpAccessor)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("main.go main() err := http.ListenAndServe() error => %v\n", err)
	}
}

/*
export GOROOT=/usr/local/go/
export GOPATH=/root/lesliehuang/goproj
export PATH=$PATH:/usr/local/go/bin/
*/
