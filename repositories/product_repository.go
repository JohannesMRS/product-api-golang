package repositories

import (
	"database/sql"
	"errors"
	"manajemen-product/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (models.Product, error)
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, created_at FROM product ORDER BY id DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	query := "INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, created_at"
	err := r.db.QueryRow(query, product.Name, product.Price).Scan(&product.ID, &product.CreatedAt)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Update(id int, product models.Product) (models.Product, error) {
	query := "UPDATE product SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price, created_at"

	err := r.db.QueryRow(query, product.Name, product.Price, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, errors.New("produk tidak ditemukan")
		}
		return product, err
	}

	return product, nil
}

func (r *productRepository) Delete(id int) error {
	query := "DELETE FROM product WHERE id = $1"

	row, err := r.db.Exec(query, id)

	affectedRow, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return errors.New("Data tidak ditemukan")
	}

	return nil
}
