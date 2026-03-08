package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/user/gocrud-api/internal/models"
)

// AccountRepository is the data access layer for accounts
type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository creates a new account repository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// GetAll returns all accounts
func (r *AccountRepository) GetAll() ([]models.Account, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.ID, &account.Name, &account.Email); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// GetByID retrieves an account by ID
func (r *AccountRepository) GetByID(id string) (*models.Account, error) {
	var account models.Account
	err := r.db.QueryRow("SELECT id, name, email FROM accounts WHERE id = $1", id).Scan(
		&account.ID, &account.Name, &account.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

// Create adds a new account
func (r *AccountRepository) Create(req models.CreateAccountRequest) (*models.Account, error) {
	id := uuid.New().String()
	account := models.Account{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}
	_, err := r.db.Exec("INSERT INTO accounts (id, name, email) VALUES ($1, $2, $3)",
		account.ID, account.Name, account.Email)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Update modifies an existing account
func (r *AccountRepository) Update(id string, req models.UpdateAccountRequest) (*models.Account, error) {
	_, err := r.db.Exec("UPDATE accounts SET name = $1, email = $2 WHERE id = $3",
		req.Name, req.Email, id)
	if err != nil {
		return nil, err
	}
	account := models.Account{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}
	return &account, nil
}

// Delete removes an account by ID
func (r *AccountRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM accounts WHERE id = $1", id)
	return err
}
