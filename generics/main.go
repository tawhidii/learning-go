package main

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (c *CustomMap[K, V]) Insert(k K, v V) error {
	c.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	map1 := NewCustomMap[string, int]()
	map1.Insert("foo", 1)

	map2 := NewCustomMap[string, float64]()
	map2.Insert("bar", 1.1)

}
