package main

import "fmt"

func main() {
	my_string := []rune("Hello, VOROTUS!")
	for i, j := 0, len(my_string)-1; i < len(my_string)/2; i, j = i+1, j-1 {
		my_string[i], my_string[j] = my_string[j], my_string[i]
	}
	fmt.Print(string(my_string))
}
