package main

import "fmt" //importing the fmt package

// hi this is the main
// funcation which has to be in every fine
func main() {
	fmt.Println("Hi this is Start")
	first()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	last()
}

func first() {
	fmt.Println("Hi this is First Steps for Go")
}

func last() {
	fmt.Println("This is end visible sadness")
}
