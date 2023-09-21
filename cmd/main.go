package main

import (
	"ISAN-api/api"
	"ISAN-api/queries"
	"fmt"
)

func main() {

	dbb, err := queries.CreateDBstruct("", ":3306")
	if err != nil {
		fmt.Println(err)
	}

	serv := api.CreateApi("127.0.0.1", ":8080", dbb)

	serv.Server.ListenAndServe()
}
