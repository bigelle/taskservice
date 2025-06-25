package schemas

type DeleteRequest struct {
	ID uint `json:"id"`
}

type DeleteResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}
