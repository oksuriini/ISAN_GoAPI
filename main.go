package main

import (
	"ISAN-api/queries"
	"fmt"
)

func main() {

	//dbb, err := queries.CreateDBstruct("", ":3306")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//isans, err := dbb.GetAll()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(isans)

	d := queries.Database{}

	fmt.Println(d.GetRowsByFilter("lol", "arvo"))
}
