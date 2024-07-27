// package main

// import "fmt"

// func BinaraySearch(arr []int, target int) bool {
// 	low, high := 0, len(arr)-1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		if target == arr[mid] {
// 			return true
// 		} else if target > arr[mid] {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}
// 	return false
// }

// func main() {
// 	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200}
// 	res := BinaraySearch(values, 70)
// 	fmt.Println(res)
// }
