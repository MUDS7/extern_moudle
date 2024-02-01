package main

import (
	"extern_moudle/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Jack")
	fmt.Println(message)
}
