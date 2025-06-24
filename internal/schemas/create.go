package schemas

type CreateRequest struct {
	TaskName       string `json:"task_name"`
	TaskDesciption string `json:"task_desciption"`
}

type CreateResponse struct {
	Ok         bool   `json:"ok"`
	TaskID     uint   `json:"task_id,omitempty"`
	TaskName   string `json:"task_name,omitempty"`
	TaskStatus string `json:"task_status,omitempty"`
	Error      string `json:"error,omitempty"`
}
