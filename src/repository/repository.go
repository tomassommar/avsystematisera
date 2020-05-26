package repository

type Repository interface {
	Stores() []Store
	Products() []Product
}

type Store interface {
	ID() string
	Name() string
	Street() string
	City() string
	PostalCode() string
	PhoneNumber() string
	Coordinates() (float64,float64)
	GetProducts() []int
}

type Product interface {
	ID() string
	Name() string
	Price() float64
	Volume() float64
	VolumeUnitPrice() float64
	AlcoholPercentage() float64
	Apk() float64
}
