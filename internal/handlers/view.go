package handlers

import (
	"errors"
	"log/slog"
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
	var status int
	db := database.NewDB()

	err = internal.ReadJSON(r, &req)
	if err != nil {
		status = http.StatusBadRequest

		resp = schemas.ViewResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, status, resp)

		slog.Info("/view", "status", http.StatusText(status), "code", status)

		return
	}

	var task database.Task
	task, err = db.View(req.ID)
	if err != nil {
		resp.Ok = false

		if errors.Is(err, database.ErrInvalidData) {
			status = http.StatusBadRequest

			resp.Error = "bad request"
			internal.WriteJSON(w, status, resp)

		} else if errors.Is(err, database.ErrNoRecord) {
			status = http.StatusNotFound

			resp.Error = "not found"
			internal.WriteJSON(w, status, resp)

		} else {
			status = http.StatusInternalServerError

			resp.Error = "internal server error"
			internal.WriteJSON(w, status, resp)
		}

		slog.Info("/view", "status", http.StatusText(status), "code", status)
		return
	}

	status = http.StatusOK

	resp = schemas.ViewResponse{
		Ok: true,
		Task: schemas.Task{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			CreatedAt:   task.CreatedAt.UTC().Format(time.RFC1123),
			UpdatedAt:   task.UpdatedAt.UTC().Format(time.RFC1123),
			Status:      task.Status.String(),
		},
	}

	if task.Status == database.StatusSuccess || task.Status == database.StatusFail {
		resp.Task.Took = time.Since(task.CreatedAt).String()
		if task.Status == database.StatusSuccess {
			resp.Task.Result = *task.Result
		}
	}

	internal.WriteJSON(w, status, resp)
	slog.Info("/view", "status", http.StatusText(status), "code", status)
}
