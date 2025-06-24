package database

import (
	"encoding/json"
	"time"
)

type Task struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Status      TaskStatus
	Result      *string
}

func NewDB() TaskDB {
	return NewLocalDB()
}

type TaskDB interface {
	Create(name, desc string) (uint, error)
	View(taskID uint) (Task, error)
	UpdateStatus(taskID uint, status string) (Task, error)
	UpdateResult(taskID uint, result string) (Task, error)
	Delete(taskID uint) error
}

type TaskStatus int

const (
	StatusUndefined = iota
	StatusPending
	StatusInProgress
	StatusCancelled
	StatusSuccess
	StatusFail
)

func (t TaskStatus) String() string {
	return [...]string{"undefined", "pending", "in_progress", "cancelled", "success", "fail"}[t]
}

func TaskStatusFromString(str string) TaskStatus {
	status, ok := map[string]TaskStatus{
		"pending":     StatusPending,
		"in_progress": StatusInProgress,
		"cancelled":   StatusCancelled,
		"success":     StatusSuccess,
		"fail":        StatusFail,
	}[str]
	if !ok {
		return StatusUndefined
	}
	return status
}

func (t TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TaskStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*t = TaskStatusFromString(str)

	return nil
}
