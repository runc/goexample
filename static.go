package main

import "net/http"

func main(){
	handler := http.FileServer(http.Dir("./"))
	http.ListenAndServe(":5000",handler)
}