package xml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type stores struct {
	XMLName   xml.Name   `xml:"ButikerOmbud"`
	Inventory []xmlStore `xml:"ButikOmbud"`
}

type products struct {
	XMLName xml.Name       `xml:"artiklar"`
	Inventory []xmlProduct `xml:"artikel"`
}

type articleMappings struct{
	XMLName xml.Name `xml:"Butik"`
	ID string `xml:"ButikNr,attr"`
	Inventory []int `xml:"ArtikelNr"`
}

type storeProducts struct{
	XMLName xml.Name            `xml:"ButikArtikel"`
	Inventory []articleMappings `xml:"Butik"`
}

func(value *alcoholPercentage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	parsed, err := strconv.ParseFloat(strings.TrimRight(s, "%"), 64)
	if err!=nil{
		return err
	}
	*value = alcoholPercentage(parsed)
	return nil
}

var storeInventory stores
var productInventory products
var productMapping storeProducts
var storeProductMapping map[string][]int
func init() {
	pwd, _ := os.Getwd()
	storeData, _ := ioutil.ReadFile(filepath.Join(pwd, "../resources/stores.xml"))
	x := xml.Unmarshal(storeData, &storeInventory)
	if x != nil {
		panic(x)
	}
	productData, _ := ioutil.ReadFile(filepath.Join(pwd, "../resources/Articles.xml"))
	x = xml.Unmarshal(productData, &productInventory)
	if x != nil {
		panic(x)
	}
	articleMappingData, _ := ioutil.ReadFile(filepath.Join(pwd, "../resources/store_articles.xml"))
	x = xml.Unmarshal(articleMappingData, &productMapping)
	if x != nil {
		panic(x)
	}
	storeProductMapping = make(map[string][]int)
	for _, articleMapping := range productMapping.Inventory{
		storeProductMapping[articleMapping.ID] = articleMapping.Inventory
	}
}
