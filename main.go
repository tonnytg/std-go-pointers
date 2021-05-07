package main

import "fmt"

func main()  {

	a := 10

	fmt.Println("Let's Go!")
	fmt.Printf("The value I have now is: %v \n", a)
	fmt.Printf("Memory of a is: %v \n", &a)

	var pointer *int = &a

	fmt.Printf("Pointer value is: %v \n", pointer)
	fmt.Println("Value of pointer is:", *pointer)

	b := 11
	pointer = &b

	fmt.Printf("Now the memory position of pointer is %v \n", &pointer)
	fmt.Printf("Now the value of a is %v and pointer is %v \n", *pointer, a)
}


