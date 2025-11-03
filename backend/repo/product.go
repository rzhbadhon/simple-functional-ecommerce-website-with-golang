package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {

	return &productRepo{
		db: db,
	}

}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
			INSERT INTO products(
		title,
		description,
		price,
		img_url
	) VALUES(
		$1,
		$2,
		$3,
		$4
	)
		RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

	query := `
		SELECT
			id,
			title,
			description,
			price,
			img_url
		from products
		where id=$1
	`
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &prd, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	var prdList []*domain.Product
	query := `
		SELECT
			id,
			title,
			description,
			price,
			img_url
		from products
	`

	err := r.db.Select(&prdList, query)

	if err != nil {
		return nil, err
	}
	return prdList, nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products SET
			title = $1,
			description = $2,
			price = $3,
			img_url = $4,
			updated_at = now()
		WHERE id = $5
		RETURNING id, title, description, price, img_url
	`
	err := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID).
		Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImgUrl)
	if err != nil {
		return nil, err
	}
	return &p, nil

}

func (r *productRepo) Delete(id int) error {
	query := `
		DELETE FROM products where id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
