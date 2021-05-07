package main

import "fmt"

func main()  {

	a := 10

	fmt.Println("Let's Go!")
	fmt.Printf("The value I have now is: %v \n", a)
	fmt.Printf("Memory of a is: %v \n", &a)

	var ponteiro *int = &a

	fmt.Printf("Valor do ponteiro é: %v \n", ponteiro)
}


