package main

/*
important note:
---------------
Capital letter -> Public Access
Small letter -> Inside that pakage acess / private access

apis folder is package



*/

import (
	"fmt"

	"tawhidii.com/apis/apis"
)

func main() {
	fmt.Println("=== Welcome to modules ==")
	apiUser := apis.UserAPIType{
		ApiVersion: 12.23,
		ApiDetails: apis.UserAPIDetails("Api for user details "),
	}
	fmt.Println(apiUser)
}
