package schemas

import "time"

type ViewRequest struct {
	TaskID uint `json:"task_id"`
}

type ViewResponse struct {
	Ok    bool   `json:"ok"`
	Task  Task   `json:"task,omitzero"`
	Error string `json:"error,omitzero"`
}

type Task struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
	Result      string    `json:"result,omitempty"`
	Took        string    `json:"took,omitempty"`
}
