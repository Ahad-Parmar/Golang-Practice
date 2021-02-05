package main

import "fmt"

func main() {
	var s [5]string
	fmt.Println(s)

	var c [5]complex128
	fmt.Println(c)

	var in [5]int

	in[0] = 100
	in[1] = 101

	fmt.Printf("in[0] = %d \n in[1] = %d\n", in[0], in[1])
	fmt.Println("in = ", in)

	var arr = [5]int{10, 20, 30, 40, 50}
	fmt.Println("arrays = ", arr)

	arr1 := [5]string{"golang", "postgre", "sql", "oracle", "badgerdb"}
	arr2 := arr1

	arr2[1] = "gittt"

	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)

	Name := [3]string{"Thor", "Hulk", "Jerry"}

	for i := 0; i < len(Name); i++ {
		fmt.Println(Name[i])
	}

	num := [4]float64{1.1, 1.2, 3.3, 3.4}
	sum := float64(0)

	for _, value := range num {
		sum = sum + value
	}

	fmt.Printf("Summing elements of array %v = %f \n", num, sum)

	mul := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(mul)

	mul2 := [3][4]float64{
		{2, 4},
		{1.5, 2.4, 3.5, -4},
		{7, 8, 6},
	}
	fmt.Println(mul2)
}
