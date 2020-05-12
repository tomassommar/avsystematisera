package main

import (
	"fmt"
	"github.com/tomassommar/avsystematisera/repository"
	"sort"
)

func main() {
	products := repository.GetProducts()
	sort.Slice(products, func(i, j int) bool {
		return products[i].Apk() < products[j].Apk()
	})
	//for _, element := range products{
	//	// fmt.Println(element.ID)
	//	fmt.Println(element.Name)
	//	fmt.Println(element.Price)
	//	fmt.Println(element.Volume)
	//
	//	fmt.Println(element.AlcoholPercentage)
	//	fmt.Println("APK: " + strconv.FormatFloat(element.Apk(), 'E', -1, 64))
	//}

	//articleMappings := repository.GetArticleMapping()
	//for _, element := range articleMappings{
	//	fmt.Println(element.ID)
	//}

	stores := repository.GetStores()
	articles := stores[0].GetArticles()
	for _, article := range articles{
		fmt.Println(article)
	}
}
