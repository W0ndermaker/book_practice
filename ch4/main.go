package main

import "fmt"

func main() {
	//a := []int{0, 1, 2, 3, 4, 5}
	//reverse(&a)
	//rotate(a)
	//fmt.Println(a)
	s := []string{"a", "b", "c", "c", "c", "d", "d"}
	dup(s)
	fmt.Println(s)
}

func dup(s []string) {
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] == s[i-1] {
			s[i] = s[len(s)-1]
		}
	}
}

/*func rotate(s []int) {
	// сдвиг на 1
	reverse(s[1:])
	reverse(s)
}

func reverse(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}*/
