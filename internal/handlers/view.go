package handlers

import (
	"net/http"
	"time"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

func HandleView(w http.ResponseWriter, r *http.Request) {
	var err error
	var req schemas.ViewRequest
	var resp schemas.ViewResponse
	db := database.NewDB()

	err = internal.ReadJSON(r, &req)
	if err != nil {
		resp = schemas.ViewResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, http.StatusBadRequest, resp)
	}

	var task database.Task
	task, err = db.View(req.TaskID)
	if err != nil {
		resp = schemas.ViewResponse{
			Ok:    false,
			Error: "internal server error",
		}
		internal.WriteJSON(w, http.StatusBadRequest, resp)
	}

	resp = schemas.ViewResponse{
		Ok: true,
		Task: schemas.Task{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			Status:      task.Status.String(),
		},
	}

	if task.Result != nil && (task.Status == database.StatusSuccess || task.Status == database.StatusFail) {
		resp.Task.Took = internal.FormatTime(time.Since(task.CreatedAt))
		if task.Status == database.StatusSuccess {
			resp.Task.Result = *task.Result
		}
	}

	internal.WriteJSON(w, http.StatusOK, resp)
}
