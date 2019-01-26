package main

import "fmt"

func listOf3() {
	list := []string{"b", "c", "a", "a", "b", "b"}

	i, left, middle := 0, 0, 0
	right := len(list) - 1

	fmt.Println(list)

	for middle <= right {
		switch list[middle] {
		case "a":
			list[left], list[middle] = list[middle], list[left]
			left++
			middle++
		case "b":
			middle++
		case "c":
			list[middle], list[right] = list[right], list[middle]
			right--
		}
		i++
	}
	fmt.Println(list)
	fmt.Println(i, "iterations")
}

func listOf2() {
	list := []int{1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1}

	i, left := 0, 0
	right := len(list) - 1

	fmt.Println(list)

	for left <= right {
		for list[left] == 0 {
			left++
			i++
		}
		for list[right] == 1 {
			right--
			i++
		}
		if left < right {
			list[left], list[right] = list[right], list[left]
			left++
			right--
			i++
		}
	}
	fmt.Println(list)
	fmt.Println(i, "iterations")
}

func main() {
	listOf2()
	listOf3()
}
