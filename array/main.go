package main

import "fmt"

func main() {
	var array [3]int

	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println(array)

	array2 := [3]string{"a", "b", "c"}

	fmt.Println(array2)

	var matrix [3][3]int

	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	fmt.Println(matrix)

	array3 := [...]int{9, 8, 7, 6, 5}

	fmt.Println(array3)

	var slice []int

	slice = []int{1, 2, 3, 4, 5}

	fmt.Println(slice)

	slice = append(slice, 99)

	fmt.Println(slice)

	arrayToRemoveASlice := [5]int{88, 89, 90, 91, 92}

	slice2 := arrayToRemoveASlice[2:4]

	fmt.Println(slice2)

	arrayToRemoveASlice[2] = 101
	arrayToRemoveASlice[3] = 102

	fmt.Println(slice2)

	slice3 := make([]float64, 5, 6)

	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
