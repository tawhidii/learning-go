// package main

// import "fmt"

// func main() {
// 	msgCh := make(chan string, 128)
// 	msgCh <- "A"
// 	msgCh <- "B"
// 	msgCh <- "C"
// 	msgCh <- "D"
// 	close(msgCh)
// 	for result := range msgCh {
// 		fmt.Println(result)
// 	}
// }
