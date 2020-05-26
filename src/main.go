package main

import (
	"fmt"
	"github.com/tomassommar/avsystematisera/repository"
	"github.com/tomassommar/avsystematisera/xml"
	"sort"
)

func main() {
	var xmlRepository repository.Repository = xml.XmlRepository{}
	products := xmlRepository.Products()
	sort.Slice(products, func(i, j int) bool {
		return products[i].Apk() < products[j].Apk()
	})

	stores := xmlRepository.Stores()
	for _, store := range stores{
		fmt.Println(store.Coordinates())
	}
	//articles := stores[0].GetProducts()
	//for _, article := range articles{
	//	fmt.Println(article)
	//}
}
