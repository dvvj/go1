package main

import (
	"fmt"
)

func main() {
	fmt.Println("Helo!")

	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	s2 := s1[1:3]

	fmt.Println("s1[1] = 22")
	s1[1] = 22
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("s1[2] = 33 after appending")
	s2 = append(s2, 4)
	s1[2] = 33
	fmt.Println(s1)
	fmt.Println(s2)

}
