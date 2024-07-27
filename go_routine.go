// package main

// import (
// 	"fmt"
// 	"time"
// )

// /*
// Notes:
// Channel in golang if it's buffered / unbuffered always blocked if it is full.
// */

// type UserData struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func main() {
// 	// user := UserData{
// 	// 	FirstName: "Khademul",
// 	// 	LastName:  "Bari",
// 	// 	Age:       225,
// 	// }
// 	//userCh := make(chan UserData) // Unbuffered
// 	// userCh := make(chan UserData, 10) // Buffered

// 	userCh := make(chan string)

// 	go func() {
// 		result := <-userCh

// 		fmt.Println(result)
// 	}()

// 	userCh <- "Hello"

// }

// func getResources() string {
// 	time.Sleep(time.Second * 2)
// 	return "Some result !"
// }
