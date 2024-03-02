package main

import (
	"fmt"
	"net/http"

	"github.com/Engineerbabu777/go-todo/database"
	"github.com/Engineerbabu777/go-todo/router"
)

func main() {

	fmt.Println("Hello From Crud Api!")

	database.DatabaseConnect();

	r := router.Router()

	// LISTENING ON THE PORT: 4444
	http.ListenAndServe(":4444", r)
}
