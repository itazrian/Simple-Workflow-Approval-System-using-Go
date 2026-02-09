package model

import "time"

type Request struct {
	ID          string `gorm:"type:char(36);primaryKey"`
	WorkflowID  string
	CurrentStep int
	Status      string
	Amount      int64
	CreatedAt   time.Time
}
