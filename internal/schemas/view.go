package schemas

type ViewRequest struct {
	ID uint `json:"id"`
}

type ViewResponse struct {
	Ok    bool   `json:"ok"`
	Task  Task   `json:"task,omitzero"`
	Error string `json:"error,omitzero"`
}

type Task struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Result      string `json:"result,omitempty"`
	Took        string `json:"took,omitempty"`
}
