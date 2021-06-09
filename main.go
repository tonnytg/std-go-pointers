package main

import (
	"errors"
	"fmt"
)
// ShowName print your name, return the same value in pointer
func ShowName(a *string) (*string, error) {
	if *a == "" {
		return a, errors.New("404")
	} else {
		return a, nil
	}
}

func main()  {
	name := "Antonio"
	valor, err := ShowName(&name)
	if err == nil {
		fmt.Printf("Name: %s\n",*valor)
	} else {
		fmt.Errorf("Error name print: %v\n", err)
	}

	name2 := ""
	valor, err = ShowName(&name2)
	if err == nil {
		fmt.Printf("Name: %s\n", *valor)
	} else {
		fmt.Errorf("Error name print: %v\n", err)
	}
}


