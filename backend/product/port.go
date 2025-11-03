package product

import (
	"ecommerce/domain"
	prdctHndlr "ecommerce/rest/handlers/product"
)

type Service interface {
	prdctHndlr.Service
}
type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productID int) error
	Update(p domain.Product) (*domain.Product, error)
}
