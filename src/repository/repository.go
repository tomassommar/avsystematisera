package repository

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetStores returns store inventory
func GetStores() []Store {
	return inventory.Inventory
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

type stores struct {
	XMLName   xml.Name `xml:"ButikerOmbud"`
	Inventory []Store  `xml:"ButikOmbud"`
}

var inventory stores

func init() {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(filepath.Join(pwd, "butiker.xml"))
	x := xml.Unmarshal(data, &inventory)
	if x != nil {
		panic(x)
	}
}
