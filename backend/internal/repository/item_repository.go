package repository

import (
	"database/sql"

	"github.com/user/gocrud-api/internal/models"
)

// ItemRepository is the data access layer for items
type ItemRepository struct {
	db *sql.DB
}

// NewItemRepository creates a new item repository
func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

// GetAll retrieves all items from the database
func (r *ItemRepository) GetAll() ([]models.Item, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, quantity, category FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Quantity, &item.Category); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// GetByID retrieves a single item by its ID
func (r *ItemRepository) GetByID(id int) (*models.Item, error) {
	var item models.Item
	err := r.db.QueryRow("SELECT id, name, description, price, quantity, category FROM items WHERE id = $1", id).Scan(
		&item.ID, &item.Name, &item.Description, &item.Price, &item.Quantity, &item.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

// Create adds a new item to the database
func (r *ItemRepository) Create(item *models.Item) (*models.Item, error) {
	err := r.db.QueryRow(
		"INSERT INTO items (name, description, price, quantity, category) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		item.Name, item.Description, item.Price, item.Quantity, item.Category).Scan(&item.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Update modifies an existing item by its ID
func (r *ItemRepository) Update(id int, item *models.Item) (*models.Item, error) {
	_, err := r.db.Exec(
		"UPDATE items SET name = $1, description = $2, price = $3, quantity = $4, category = $5 WHERE id = $6",
		item.Name, item.Description, item.Price, item.Quantity, item.Category, id)
	if err != nil {
		return nil, err
	}
	item.ID = id
	return item, nil
}

// Delete removes an item by its ID
func (r *ItemRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}

// GetByCategory retrieves all items belonging to a specific category
func (r *ItemRepository) GetByCategory(category string) ([]models.Item, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, quantity, category FROM items WHERE category = $1", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Quantity, &item.Category); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Search performs a case-insensitive search on item names
func (r *ItemRepository) Search(query string) ([]models.Item, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, quantity, category FROM items WHERE LOWER(name) LIKE LOWER($1)", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Quantity, &item.Category); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
