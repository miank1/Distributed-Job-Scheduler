package scheduler

import "time"

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusDone    Status = "done"
	StatusFailed  Status = "failed"
)

type Job struct {
	ID         string
	TaskType   string
	Payload    []byte
	RunAt      time.Time
	Status     Status
	RetryCount int
	MaxRetries int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LastError  string
}
