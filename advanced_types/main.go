package main

import "fmt"

type User struct {
	username string
	password string
}

type Profile struct {
	User
	firstName string
	lastName  string
}

func main() {
	profile := Profile{
		User: User{
			username: "testuser",
			password: "abc123abc",
		},
		firstName: "Khademul",
		lastName:  "Bari",
	}
	fmt.Println(profile)
}

// Another Example 
package main

import "fmt"

type Color int

const (
	ColorBlack Color = iota
	ColorBlue
	ColorRed
	ColorGreen
)

// fmt Stringer
func (c Color) String() string {
	switch c {
	case ColorGreen:
		return "Green"
	case ColorBlue:
		return "Blue"
	case ColorRed:
		return "Red"
	case ColorBlack:
		return "Black"
	default:
		panic("In valid color provided")
	}

}

func main() {
	fmt.Println(ColorGreen)
}
