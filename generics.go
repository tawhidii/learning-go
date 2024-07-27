// package main

// import "fmt"

// type CustomMap[K comparable, V any] struct {
// 	data map[K]V
// }

// func (m *CustomMap[K, V]) Insert(key K, value V) error {
// 	m.data[key] = value
// 	return nil
// }

// func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
// 	return &CustomMap[K, V]{data: make(map[K]V)}
// }

// // function generics

// func Something[N any, A any](name N, age A) {
// 	fmt.Println(name, age)
// }

// func main() {
// 	m1 := NewCustomMap[string, int]()
// 	m1.Insert("Hello", 100)
// 	Something[string, string]("Bari", "100")
// }
