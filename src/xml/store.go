package xml

import (
	"encoding/xml"
	"github.com/tomassommar/avsystematisera/repository"
	"github.com/tomassommar/swedishgrid"
	"strconv"
)

type xmlStore struct {
	XMLName        xml.Name `xml:"ButikOmbud"`
	XmlId          string   `xml:"Nr"`
	XmlName        string   `xml:"Namn"`
	XmlStreet      string   `xml:"Address1"`
	XmlCity        string   `xml:"Address4"`
	XmlPostalCode  string   `xml:"Address3"`
	XmlPhoneNumber string   `xml:"Telefon"`
	XmlRT90x          string   `xml:"RT90x"`
	XmlRT90y          string   `xml:"RT90y"`
}

func (s stores) Stores() []repository.Store{
	stores := make([]repository.Store, 0)
	for _,v:=range s.Inventory{
		stores = append(stores, v)
	}
	return stores
}

func (store xmlStore) ID() string {
	return store.XmlId
}

func (store xmlStore) Name() string {
	return store.XmlName
}

func (store xmlStore) Street() string {
	return store.XmlStreet
}

func (store xmlStore) City() string {
	return store.XmlCity
}

func (store xmlStore) PostalCode() string {
	return store.XmlPostalCode
}

func (store xmlStore) PhoneNumber() string {
	return store.XmlPhoneNumber
}

func (store xmlStore) Coordinates() (float64,float64) {
	x, err := strconv.ParseFloat(store.XmlRT90x, 64)
	y, err := strconv.ParseFloat(store.XmlRT90y, 64)
	if err!=nil{
		panic(err)
	}
	return swedishgrid.RT90toWGS84(x, y)
}
