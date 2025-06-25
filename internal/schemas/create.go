package schemas

type CreateRequest struct {
	TaskName       string `json:"name"`
	TaskDesciption string `json:"description"`
}

type CreateResponse struct {
	Ok     bool   `json:"ok"`
	ID     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}
