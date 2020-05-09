package main

import (
	"fmt"
	"github.com/tomassommar/avsystematisera/repository"
)

func main() {
	stores := repository.GetStores()
	for _, element := range stores {
		fmt.Println(element.ID)
		fmt.Println(element.Name)
		fmt.Println(element.Street)
		fmt.Println(element.City)
		fmt.Println(element.PostalCode)
	}

}
