package main

import (
	"fmt"
)

// types
var (
	name   string  = "Harry potter"
	height float32 = 5.6
	wieght int64   = 80
	//... many more
)

type Player struct {
	name        string
	height      float32
	wieght      int32
	nationality string
}

// passing Player struct as argument
func getPlayerName(player Player) string {
	return player.name
}

// example of reciever function
func (p Player) getPlayerNationality() string {
	return p.nationality
}

// custom types
type Person string

func getPerson(person Person) Person {
	return person
}

// main func
func main() {
	playerOne := Player{} // will initial with default value
	playertwo := Player{
		name:        "Bari",
		height:      5.6,
		wieght:      80,
		nationality: "Bangladeshi",
	}
	person := getPerson("Bari")
	fmt.Println(person)

	fmt.Printf("The player 1 is %v", playerOne)  // Printf -> Formatted print
	fmt.Printf("The player 2 is %+v", playertwo) // +v is for super verbose mode
	fmt.Println("Player 2 name is ", getPlayerName(playertwo))

	fmt.Println("Player 2 nationality is ", playertwo.getPlayerNationality())

	// example of map
	users := map[string]int64{
		"age": 25,
	}
	// another way to declare map
	otherUsers := make(map[string]int64)
	fmt.Println(users)
	otherUsers["salary"] = 0
	otherUsers["siblings"] = 4
	fmt.Println(otherUsers)

	siblings, exists := otherUsers["siblings"] // check key exsists or not in go
	if !exists {
		fmt.Println("siblings key doesnot exists")
	} else {
		fmt.Println("siblings exists", siblings)
	}

	// array
	fruits := [3]string{"apple", "orange"} // you can not append to array ,it works on slice
	fmt.Println(fruits)

	// slice
	movies := []string{}
	movies = append(movies, "Harry potter")
	movies = append(movies, "Batman")
	fmt.Println(movies)

}
