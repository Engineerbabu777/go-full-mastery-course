package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request){
   fmt.Fprint(w,"Hello From Home route!")
}


func main(){

	http.HandleFunc("/", handler)
	http.ListenAndServe(":7544", nil)
}