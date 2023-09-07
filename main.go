package main

import (
	"ISAN-api/queries"
	"fmt"
)

func main() {

	dbb, err := queries.CreateDBstruct("", ":3306")
	if err != nil {
		fmt.Println(err)
	}

	isans, err := dbb.GetRowsByFilter("work_typ", "120")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(isans)

}
