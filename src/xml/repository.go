package xml

import "github.com/tomassommar/avsystematisera/repository"

type XmlRepository struct {}

func(repository XmlRepository) Products() []repository.Product {
	return productInventory.Products()
}

func(repository XmlRepository) Stores() []repository.Store {
	return storeInventory.Stores()
}

func (store xmlStore) GetArticles() []int{
	return storeProductMapping[store.XmlId]
}
