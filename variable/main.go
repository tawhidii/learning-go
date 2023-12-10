package main

// global variables

var firstName = "Bari"
var lastName string = "Tawhidi"
var gender string

// Grouping global variable
var (
	fatherName string = "ABC"
	motherName        = "XYZ"
	address    string
)

// declaring constant variable
const pi float32 = 3.1416

// Group decleration const variable
const (
	constVar  string  = "ABC"
	constVar2 float64 = 3.34444
)

func main() {
	// local variables
	version := 10                  // infer int
	newVersion := 20.45            //infer float
	anotherNewVersion := "version" // infer string

}
