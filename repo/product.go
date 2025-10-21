package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"` // = it is called tag
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
	ImageURL    string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Get(productID int) (*Product, error)
	Create(product Product) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(productID int) error
	List() ([]*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

// Constructor or Constructor function. It creates an object of struct.
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Get(productID int) (*Product, error) {
	var product Product

	query := `
	SELECT id, title, description, price, img_url 
	FROM product 
	WHERE id = $1
	`
	err := r.db.Get(&product, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return &product, err
}

func (r *productRepo) Create(product Product) (*Product, error) {
	query := `
	INSERT INTO product (
	title,
	description,
	price,
	img_url
	)
	VALUES(
	$1,
	$2,
	$3,
	$4
	)
	RETURNING id
	`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageURL)
	err := row.Scan(&product.ID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	query := `
	UPDATE product
	SET title =$1, description =$2, price =$3, img_url =$4
	WHERE id = $5
	`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageURL)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) Delete(productID int) error {
	query := `
	DELETE FROM product WHERE id = $1
	`
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product

	query := `
		SELECT id, title, description, price, img_url
		FROM product
	`
	err := r.db.Select(&productList, query)
	if err != nil {
		return nil, err
	}

	return productList, nil
}
