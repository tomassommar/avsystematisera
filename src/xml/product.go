package xml

import "github.com/tomassommar/avsystematisera/repository"

type xmlProduct struct {
	XmlId                string            `xml:"Artikelid"`
	XmlName              string            `xml:"Namn"`
	XmlPrice             float64           `xml:"Prisinklmoms"`
	XmlVolume            float64           `xml:"Volymiml"`
	XmlVolumeUnitPrice   float64           `xml:"PrisPerLiter"`
	XmlAlcoholPercentage alcoholPercentage `xml:"Alkoholhalt"`
}

type alcoholPercentage float64

func (a products) Products() []repository.Product{
	articles := make([]repository.Product, 0)
	for _,v:=range a.Inventory{
		articles = append(articles, v)
	}
	return articles
}

func (product xmlProduct) ID() string {
	return product.XmlId
}

func (product xmlProduct) Name() string {
	return product.XmlName
}

func (product xmlProduct) Price() float64 {
	return product.XmlPrice
}

func (product xmlProduct) Volume() float64 {
	return product.XmlVolume
}

func (product xmlProduct) VolumeUnitPrice() float64 {
	return product.XmlVolumeUnitPrice
}

func (product xmlProduct) AlcoholPercentage() float64 {
	return float64(product.XmlAlcoholPercentage)
}

func (product xmlProduct) Apk() float64{
	alcoholVolume := product.XmlVolume * (float64(product.XmlAlcoholPercentage)/100)
	return alcoholVolume/product.XmlPrice
}
