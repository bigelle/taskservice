package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

type CreateHandler struct{}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	db := database.NewDB()
	dec := internal.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	var req schemas.CreateRequest
	var resp schemas.CreateResponse

	err = dec.Decode(&req)
	if err != nil {
		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "bad request",
		}
		w.WriteHeader(http.StatusBadRequest)
		enc.Encode(resp)
	}

	var taskID uint
	taskID, err = db.Create(
		req.CreatorName,
		req.TaskName,
		req.TaskDesciption,
	)
	if err != nil {
		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "internal server error",
		}
		w.WriteHeader(http.StatusInternalServerError)
		enc.Encode(resp)
	}

	resp = schemas.CreateResponse{
		Ok:       true,
		TaskName: req.TaskName,
		TaskID:   taskID,
	}
	w.WriteHeader(http.StatusOK)
	enc.Encode(resp)
}
