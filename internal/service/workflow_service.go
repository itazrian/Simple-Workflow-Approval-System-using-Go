package service

import (
	"errors"
	"time"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowService struct {
	db *gorm.DB
}

func NewWorkflowService(db *gorm.DB) *WorkflowService {
	return &WorkflowService{db: db}
}

var workflows = make(map[string]model.Workflow)

func (s *WorkflowService) CreateWorkflow(wf *model.Workflow) (*model.Workflow, error) {
	if wf.Name == "" {
		return nil, errors.New("workflow name is required")
	}
	wf.ID = uuid.New().String()
	wf.CreatedAt = time.Now()
	if err := s.db.Create(wf).Error; err != nil {
		return nil, err
	}
	return wf, nil
}

func (s *WorkflowService) GetAllWorkflows() ([]model.Workflow, error) {
	var workflows []model.Workflow
	if err := s.db.Find(&workflows).Error; err != nil {
		return nil, err
	}
	return workflows, nil
}

func (s *WorkflowService) GetWorkflowByID(id string) (*model.Workflow, error) {
	var wf model.Workflow
	if err := s.db.First(&wf, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &wf, nil
}
