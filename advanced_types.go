// // package main

// // import "fmt"

// // // Example of Embedding Types
// // type SuperHero struct {
// // 	Name  string
// // 	Power string
// // }

// // type Comic struct {
// // 	SuperHero
// // 	ComicName string
// // }

// // func main() {

// // 	comic := Comic{ComicName: "Marvel", SuperHero: SuperHero{
// // 		Name: "Iron Man", Power: "Has Machines",
// // 	}}
// // 	comic.ComicName = "DC"
// // 	fmt.Println(comic.ComicName)

// // }

// package main

// import "fmt"

// // Example of Enums

// type SuperHero int

// const (
// 	SuperMan SuperHero = iota
// 	SpiderMan
// 	IronMan
// 	AntMan
// 	AquaMan
// 	BatMan
// )

// // fmt.Stringer
// func (s SuperHero) String() string {
// 	switch s {
// 	case SuperMan:
// 		return "FLY"
// 	case SpiderMan:
// 		return "WEB SHOOTER"
// 	case IronMan:
// 		return "NANO TECH"
// 	case AntMan:
// 		return "CAN SMALL"
// 	case AquaMan:
// 		return "RULES OCEAN"
// 	case BatMan:
// 		return "BAT MOBILE"
// 	default:
// 		panic("Nothing found !")
// 	}
// }

// func main() {

// 	fmt.Println(BatMan)

// }
