package main

import "net/url"

func main() {

	const URL string = "https://localhost:4444/pass?re=nane"

	result, err := url.Parse(URL)

	if err != nil {
		panic(err)
	}

	// result.Schema
	// result.HOST
	// result.Query();
	// result.rawQuery;
	// result.path

}
