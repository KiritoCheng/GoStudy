package main

import "fmt"

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In moddify() , array values:", array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(),arrray values:", array)

	/*******数组切片array[first:last]********/
	fmt.Println("\n/*******数组切片array[first:last]********/\n")

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Elemets of myArray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElemets of mySlice:")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	/*******数组切片append()********/
	fmt.Println("\n/*******数组切片append()********/\n")

	mySlice2 := make([]int, 5, 10)

	fmt.Println("len(mySlice2)：", len(mySlice2))
	fmt.Println("cap(mySlice2)：", cap(mySlice2))

	mySlice2 = append(mySlice2, mySlice...)
	fmt.Println(mySlice2)

	/*******数组切片copy()********/
	fmt.Println("\n/*******数组切片copy()********/\n")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	fmt.Println(slice1, slice2)

	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	copy(slice1, slice2)
	fmt.Println(slice1, slice2)
}
