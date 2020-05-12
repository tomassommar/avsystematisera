package repository

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GetStores returns store storeInventory
func GetStores() []Store {
	return storeInventory.Inventory
}

func GetProducts() []Product{
	return productInventory.Inventory
}

func GetArticleMapping() []ArticleMapping{
	return articleMapping.Inventory
}

// Store contains ID, address and coordinates on RT90x format
type Store struct {
	XMLName     xml.Name `xml:"ButikOmbud"`
	ID          string   `xml:"Nr"`
	Name        string   `xml:"Namn"`
	Street      string   `xml:"Address1"`
	City        string   `xml:"Address4"`
	PostalCode  string   `xml:"Address3"`
	PhoneNumber string   `xml:"Telefon"`
	RT90x       string   `xml:"RT90x"`
	RT90y       string   `xmö:"RT90y"`
}

type AlcoholPercentage float64
func(value *AlcoholPercentage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	parsed, err := strconv.ParseFloat(strings.TrimRight(s, "%"), 64)
	if err!=nil{
		return err
	}
	*value = AlcoholPercentage(parsed)
	return nil
}

type Product struct {
	ID string `xml:"Artikelid"`
	Name string `xml:"Namn"`
	Price float64 `xml:"Prisinklmoms"`
	Volume float64 `xml:"Volymiml"`
	VolumeUnitPrice float64 `xml:"PrisPerLiter"`
	AlcoholPercentage AlcoholPercentage `xml:"Alkoholhalt"`
}

func (product Product) Apk() float64{
	alcoholVolume := product.Volume * (float64(product.AlcoholPercentage)/100)
	return alcoholVolume/product.Price
}

func (store Store) GetArticles() []int{
	return storeArticleMapping[store.ID]
}

type stores struct {
	XMLName   xml.Name `xml:"ButikerOmbud"`
	Inventory []Store  `xml:"ButikOmbud"`
}

type articles struct {
	XMLName xml.Name `xml:"artiklar"`
	Inventory []Product `xml:"artikel"`
}

type ArticleMapping struct{
	XMLName xml.Name `xml:"Butik"`
	ID string `xml:"ButikNr,attr"`
	Inventory []int `xml:"ArtikelNr"`
}

type storeArticles struct{
	XMLName xml.Name `xml:"ButikArtikel"`
	Inventory []ArticleMapping `xml:"Butik"`
}

var storeInventory stores
var productInventory articles
var articleMapping storeArticles
var storeArticleMapping map[string][]int
func init() {
	pwd, _ := os.Getwd()
	storeData, _ := ioutil.ReadFile(filepath.Join(pwd, "../resources/stores.xml"))
	x := xml.Unmarshal(storeData, &storeInventory)
	if x != nil {
		panic(x)
	}
	productData, _ := ioutil.ReadFile(filepath.Join(pwd, "../resources/articles.xml"))
	x = xml.Unmarshal(productData, &productInventory)
	if x != nil {
		panic(x)
	}
	articleMappingData, _ := ioutil.ReadFile(filepath.Join(pwd, "../resources/store_articles.xml"))
	x = xml.Unmarshal(articleMappingData, &articleMapping)
	if x != nil {
		panic(x)
	}
	storeArticleMapping = make(map[string][]int)
	for _, articleMapping := range articleMapping.Inventory{
		storeArticleMapping[articleMapping.ID] = articleMapping.Inventory
	}
}
