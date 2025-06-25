package handlers

import (
	"errors"
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
		return
	}

	var task database.Task
	task, err = db.View(req.TaskID)
	if err != nil {
		resp.Ok = false

		if errors.Is(err, database.ErrInvalidData) {
			resp.Error = "bad request"
			internal.WriteJSON(w, http.StatusBadRequest, resp)
			return
		} else if errors.Is(err, database.ErrNoRecord) {
			resp.Error = "not found"
			internal.WriteJSON(w, http.StatusNotFound, resp)
			return
		} else {
			resp.Error = "internal server error"
			internal.WriteJSON(w, http.StatusInternalServerError, resp)
			return
		}
	}

	resp = schemas.ViewResponse{
		Ok: true,
		Task: schemas.Task{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			Status:      task.Status.String(),
		},
	}

	if task.Status == database.StatusSuccess || task.Status == database.StatusFail {
		resp.Task.Took = internal.FormatTime(time.Since(task.CreatedAt))
		if task.Status == database.StatusSuccess {
			resp.Task.Result = *task.Result
		}
	}

	internal.WriteJSON(w, http.StatusOK, resp)
}
