package schemas

type UpdateRequest struct {
	TaskID    uint   `json:"task_id"`
	NewStatus string `json:"new_status"`
	Result    string `json:"result,omitempty"`
}

type UpdateResponse struct {
	Ok     bool   `json:"ok"`
	TaskID uint   `json:"task_id,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}
