package handlers

import (
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	var req schemas.UpdateRequest
	var resp schemas.UpdateResponse
	db := database.NewDB()

	err = internal.ReadJSON(r, &req)
	if err != nil {
		resp = schemas.UpdateResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, http.StatusBadRequest, resp)
		return
	}

	var task database.Task
	task, err = db.UpdateStatus(req.TaskID, req.NewStatus)
	if err != nil {
		resp = schemas.UpdateResponse{
			Ok:    false,
			Error: "internal server error",
		}
		internal.WriteJSON(w, http.StatusInternalServerError, resp)
		return
	}

	if req.Result != "" {
		task, err = db.UpdateResult(req.TaskID, req.Result)
		if err != nil {
			resp = schemas.UpdateResponse{
				Ok:    false,
				Error: "internal server error",
			}
			internal.WriteJSON(w, http.StatusInternalServerError, resp)
			return
		}
	}

	resp = schemas.UpdateResponse{
		Ok:     true,
		TaskID: task.ID,
		Status: task.Status.String(),
	}
	internal.WriteJSON(w, http.StatusOK, resp)
}

