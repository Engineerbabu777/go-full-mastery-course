package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:3000"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	dataBytes, _ := ioutil.ReadAll(response.Body)

	content := string(dataBytes)

	fmt.Println(content)

}
