package main

import "net/http"

func hello(res http.ResponseWriter,req *http.Request){
	// res.Header().Set("Content-Type","text/plain")
	res.Header().Set("Content-Type","text/plain")
	res.Write([]byte("hello web\n"))
}

func main(){
	http.HandleFunc("/",hello)
	http.ListenAndServe(":5000",nil)
}

// import "net/http"

// func hello(res http.ResponseWriter, req *http.Request) {
//     res.Header().Set("Content-Type", "text/plain")
//     res.Write([]byte("Hello web\n"))
// }

// func main() {
//     http.HandleFunc("/", hello)
//     http.ListenAndServe(":5000", nil)
// }