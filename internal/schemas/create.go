package schemas

type CreateRequest struct {
	CreatorName    string `json:"creator_name"`
	TaskName       string `json:"task_name"`
	TaskDesciption string `json:"task_desciption"`
}

type CreateResponse struct {
	Ok       bool   `json:"ok"`
	TaskID   uint   `json:"task_id,omitempty"`
	TaskName string `json:"task_name,omitempty"`
	Error    string `json:"error,omitempty"`
}
