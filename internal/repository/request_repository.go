package repository

import (
	"workflow-approval/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RequestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) *RequestRepository {
	return &RequestRepository{db}
}

func (r *RequestRepository) FindForUpdate(tx *gorm.DB, id string) (*model.Request, error) {
	var req model.Request
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&req, "id = ?", id).Error
	return &req, err
}

func (r *RequestRepository) Save(tx *gorm.DB, req *model.Request) error {
	return tx.Save(req).Error
}
