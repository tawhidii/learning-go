// package main

// import (
// 	"fmt"
// 	"time"
// )

// type Server struct {
// 	quitCh chan struct{}
// 	msgCh  chan string
// }

// func newServer() *Server {
// 	return &Server{
// 		quitCh: make(chan struct{}),
// 		msgCh:  make(chan string, 128),
// 	}
// }

// func (s *Server) start() {
// 	fmt.Println("Starting Server !!")
// 	s.loop()
// }

// func (s *Server) loop() {

// 	for {
// 		select {
// 		case <-s.quitCh:
// 		// do something to quit channel
// 		case msg := <-s.msgCh:
// 			s.handleMessage(msg)
// 		default:
// 			// do something
// 		}

// 	}
// }
// func (s *Server) handleMessage(msg string) {
// 	fmt.Println("We recieved a message", msg)
// }

// func (s *Server) sendMessage(msg string) {
// 	s.msgCh <- msg
// }

// func main() {
// 	server := newServer()
// 	go server.start()
// 	for i:= 0, i<=100, i++ {
// 		server.sendMessage(fmt.Sprintf("Number is %d",i))
// 	}

// 	time.Sleep(time.Second * 2)
// }
