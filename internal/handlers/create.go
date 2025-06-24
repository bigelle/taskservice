package handlers

import (
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	db := database.NewDB()
	var resp schemas.CreateResponse

	var req schemas.CreateRequest
	err = internal.ReadJSON(r, &req)
	if err != nil {
		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, http.StatusBadRequest, resp)
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
		internal.WriteJSON(w, http.StatusInternalServerError, resp)
	}

	resp = schemas.CreateResponse{
		Ok:       true,
		TaskName: req.TaskName,
		TaskID:   taskID,
	}
	internal.WriteJSON(w, http.StatusOK, resp)
}
