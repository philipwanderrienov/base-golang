package repository

import (
	"fmt"
	"sync"

	"github.com/user/gocrud-api/internal/models"
)

// AccountRepository is the data access layer for accounts
// In a real application, this would interface with a database
// Here we use an in-memory slice with thread-safe operations using mutex
type AccountRepository struct {
	accounts []models.Account
	mu       sync.RWMutex
	nextID   int
}

// NewAccountRepository creates a new account repository with dummy data
func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		accounts: []models.Account{
			{ID: "1", Name: "Alice Smith", Email: "alice@example.com"},
			{ID: "2", Name: "Bob Johnson", Email: "bob@example.com"},
		},
		nextID: 3,
	}
}

// GetAll returns all accounts
func (r *AccountRepository) GetAll() []models.Account {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.accounts
}

// GetByID retrieves an account by ID
func (r *AccountRepository) GetByID(id string) (*models.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, account := range r.accounts {
		if account.ID == id {
			return &account, nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

// Create adds a new account
func (r *AccountRepository) Create(req models.CreateAccountRequest) (*models.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	account := models.Account{
		ID:    fmt.Sprintf("%d", r.nextID),
		Name:  req.Name,
		Email: req.Email,
	}
	r.nextID++
	r.accounts = append(r.accounts, account)
	return &account, nil
}

// Update modifies an existing account
func (r *AccountRepository) Update(id string, req models.UpdateAccountRequest) (*models.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, account := range r.accounts {
		if account.ID == id {
			updated := models.Account{
				ID:    id,
				Name:  req.Name,
				Email: req.Email,
			}
			r.accounts[i] = updated
			return &updated, nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

// Delete removes an account by ID
func (r *AccountRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, account := range r.accounts {
		if account.ID == id {
			r.accounts = append(r.accounts[:i], r.accounts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("account not found")
}
