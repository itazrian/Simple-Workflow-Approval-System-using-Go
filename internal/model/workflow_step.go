package model

type WorkflowStep struct {
	ID         string `gorm:"type:char(36);primaryKey"`
	WorkflowID string
	Level      int
	Actor      string
	MinAmount  int64
}
