package main

import (
	"fmt"
	"rsc.io/quote"
	"test-app/user"
)

func main() {

	fmt.Println("Hello, ", "Avijit Bairagi")
	fmt.Println(quote.Go())
	fmt.Println(Get())
	fmt.Println(user.Get())
}
