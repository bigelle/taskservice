package schemas

type DeleteRequest struct {
	TaskID uint `json:"task_id"`
}

type DeleteResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}
