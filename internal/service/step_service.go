package service

import (
	"errors"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StepService struct {
	db *gorm.DB
}

func NewStepService(db *gorm.DB) *StepService {
	return &StepService{db: db}
}

func (s *StepService) CreateStep(step *model.WorkflowStep) (*model.WorkflowStep, error) {
	if step.Level <= 0 {
		return nil, errors.New("level must be greater than 0")
	}
	if step.Actor == "" {
		return nil, errors.New("actor is required")
	}
	step.ID = uuid.New().String()
	if err := s.db.Create(step).Error; err != nil {
		return nil, err
	}
	return step, nil
}

func (s *StepService) GetSteps(workflowID string) ([]model.WorkflowStep, error) {
	var steps []model.WorkflowStep
	if err := s.db.Where("workflow_id = ?", workflowID).Order("level asc").Find(&steps).Error; err != nil {
		return nil, err
	}
	return steps, nil
}
