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
	Update(taskID uint, task Task) (Task, error)
	Delete(taskID uint) error
}

type TaskStatus int

const (
	StatusPending = iota
	StatusInProgress
	StatusCancelled
	StatusSuccess
	StatusFail
)

func (t TaskStatus) String() string {
	return [...]string{"pending", "in_progress", "cancelled", "success", "fail"}[t]
}

func (t *TaskStatus) FromString(str string) TaskStatus {
	return map[string]TaskStatus{
		"pending":     StatusPending,
		"in_progress": StatusInProgress,
		"cancelled":   StatusCancelled,
		"success":     StatusSuccess,
		"fail":        StatusFail,
	}[str]
}

func (t TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TaskStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*t = t.FromString(str)

	return nil
}
