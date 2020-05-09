package main

import (
	"fmt"
	"repository"
)

func main() {
	stores := repository.GetStores()
	for _, element := range stores {
		fmt.Println(string(element.ID))
		fmt.Println(string(element.Name))
		fmt.Println(string(element.Street))
		fmt.Println(string(element.City))
		fmt.Println(string(element.PostalCode))
	}

}
