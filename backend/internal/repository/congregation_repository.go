package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/user/gocrud-api/internal/models"
)

// CongregationRepository is the data access layer for congregations
type CongregationRepository struct {
	db *sql.DB
}

// NewCongregationRepository creates a new congregation repository
func NewCongregationRepository(db *sql.DB) *CongregationRepository {
	return &CongregationRepository{db: db}
}

// GetAll returns all congregations
func (r *CongregationRepository) GetAll() ([]models.Congregation, error) {
	rows, err := r.db.Query("SELECT id, name, location FROM congregations")
	if err != nil {
		// log the SQL error for troubleshooting
		log.Printf("error querying congregations: %v", err)
		return nil, err
	}
	defer rows.Close()

	var congregations []models.Congregation
	for rows.Next() {
		var congregation models.Congregation
		if err := rows.Scan(&congregation.ID, &congregation.Name, &congregation.Location); err != nil {
			return nil, err
		}
		congregations = append(congregations, congregation)
	}
	return congregations, nil
}

// GetByID retrieves a congregation by ID
func (r *CongregationRepository) GetByID(id string) (*models.Congregation, error) {
	var congregation models.Congregation
	err := r.db.QueryRow("SELECT id, name, location FROM congregations WHERE id = $1", id).Scan(
		&congregation.ID, &congregation.Name, &congregation.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &congregation, nil
}

// Create adds a new congregation
func (r *CongregationRepository) Create(req models.CreateCongregationRequest) (*models.Congregation, error) {
	id := uuid.New().String()
	congregation := models.Congregation{
		ID:       id,
		Name:     req.Name,
		Location: req.Location,
	}
	_, err := r.db.Exec("INSERT INTO congregations (id, name, location) VALUES ($1, $2, $3)",
		congregation.ID, congregation.Name, congregation.Location)
	if err != nil {
		return nil, err
	}
	return &congregation, nil
}

// Update modifies an existing congregation
func (r *CongregationRepository) Update(id string, req models.UpdateCongregationRequest) (*models.Congregation, error) {
	_, err := r.db.Exec("UPDATE congregations SET name = $1, location = $2 WHERE id = $3",
		req.Name, req.Location, id)
	if err != nil {
		return nil, err
	}
	congregation := models.Congregation{
		ID:       id,
		Name:     req.Name,
		Location: req.Location,
	}
	return &congregation, nil
}

// Delete removes a congregation by ID
func (r *CongregationRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM congregations WHERE id = $1", id)
	return err
}
