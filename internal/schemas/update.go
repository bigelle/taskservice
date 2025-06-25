package schemas

type UpdateRequest struct {
	ID        uint   `json:"id"`
	NewStatus string `json:"new_status"`
	Result    string `json:"result,omitempty"`
}

type UpdateResponse struct {
	Ok     bool   `json:"ok"`
	ID     uint   `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}
