package main

import (
	"fmt"

	database "local.com/rest/db/orm/dbconnection"
)

func main() {
	fmt.Println("Go ORM Tutorial")
	database.InitDb()
	handleRequests()
}
