package service

import (
	"errors"
	"time"

	"workflow-approval/internal/model"
	"workflow-approval/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RequestService struct {
	db   *gorm.DB
	repo *repository.RequestRepository
}

func NewRequestService(db *gorm.DB, repo *repository.RequestRepository) *RequestService {
	return &RequestService{db, repo}
}

func (s *RequestService) CreateRequest(r *model.Request) (*model.Request, error) {
	if r.Amount <= 0 {
		return nil, errors.New("amount must be > 0")
	}
	r.ID = uuid.New().String()
	r.Status = "PENDING"
	r.CurrentStep = 1
	r.CreatedAt = time.Now()
	if err := s.db.Create(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (s *RequestService) GetRequestByID(id string) (*model.Request, error) {
	var r model.Request
	if err := s.db.First(&r, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *RequestService) Approve(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		req, err := s.repo.FindForUpdate(tx, id)
		if err != nil {
			return err
		}

		if req.Status != "PENDING" {
			return errors.New("request already processed")
		}

		var nextStep model.WorkflowStep
		err = tx.Where(
			"workflow_id = ? AND level = ? AND min_amount <= ?",
			req.WorkflowID,
			req.CurrentStep+1,
			req.Amount,
		).First(&nextStep).Error

		if err != nil {
			req.Status = "APPROVED"
		} else {
			req.CurrentStep++
		}

		return s.repo.Save(tx, req)
	})
}

func (s *RequestService) Reject(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		req, err := s.repo.FindForUpdate(tx, id)
		if err != nil {
			return err
		}

		if req.Status != "PENDING" {
			return errors.New("request already processed")
		}

		req.Status = "REJECTED"
		return s.repo.Save(tx, req)
	})
}
