package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type RepositoryProductStoreDb struct {
	db *sql.DB
}

func NewRepositoryProductStoreDb(db *sql.DB) *RepositoryProductStoreDb {
	return &RepositoryProductStoreDb{
		db: db,
	}
}

func (r *RepositoryProductStoreDb) FindById(id int) (p internal.Product, err error) {
	query := "SELECT * FROM products p WHERE p.Id = ?"
	row := r.db.QueryRow(query, id)

	var product internal.Product
	var expirationStr string
	err = row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue,
		&product.IsPublished, &expirationStr, &product.Price, &product.IdWareHouse)
	if err != nil {
		return internal.Product{}, err
	}

	t, err := time.Parse(time.DateOnly, expirationStr)
	if err != nil {
		return internal.Product{}, err
	}
	product.Expiration = t

	return product, nil
}

func (r *RepositoryProductStoreDb) Save(p *internal.Product) (err error) {
	query := "INSERT INTO products(name, quantity, code_value, is_published, expiration, price, id_warehouse) VALUES(?, ?, ?, ?, ?, ?, ?);"
	row, err := r.db.Exec(query, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price, &p.IdWareHouse)
	if err != nil {
		return err
	}

	_, err = row.LastInsertId()
	if err != nil {
		log.Printf("Atenção: não foi possível obter o ID do último inserido: %v", err)
	}

	_, err = row.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %w", err)
	}

	return nil
}

func (r *RepositoryProductStoreDb) UpdateOrSave(p *internal.Product) (err error) {
	query := "UPDATE products p SET p.name = ?, p.quantity = ?, p.code_value = ?, p.is_published = ?, p.expiration = ?, p.price = ?, p.id_warehouse = ? WHERE p.code_value = ?;"
	row, err := r.db.Exec(query, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price, &p.IdWareHouse, &p.CodeValue)
	if err != nil {
		return err
	}

	_, err = row.LastInsertId()
	if err != nil {
		log.Printf("Atenção: não foi possível obter o ID do último inserido: %v", err)
	}

	_, err = row.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %w", err)
	}

	return nil
}

func (r *RepositoryProductStoreDb) Update(p *internal.Product) (err error) {
	return nil
}

func (r *RepositoryProductStoreDb) Delete(id int) (err error) {
	query := "DELETE FROM products p WHERE p.Id = ?"
	row, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	_, err = row.LastInsertId()
	if err != nil {
		log.Printf("Atenção: não foi possível obter o ID do último inserido: %v", err)
	}

	_, err = row.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas: %w", err)
	}

	return nil
}
