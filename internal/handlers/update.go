package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

// Handles /update endpoint
func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		req    schemas.UpdateRequest
		resp   schemas.UpdateResponse
		status int
	)
	db := database.NewDB()

	err = internal.ReadJSON(r, &req)
	if err != nil {
		status = http.StatusBadRequest

		resp = schemas.UpdateResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, status, resp)

		slog.Info("/update", "status", http.StatusText(status), "code", status)
		return
	}

	var task database.Task
	task, err = db.UpdateStatus(req.ID, req.NewStatus)
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
			status = http.StatusNotFound

			resp.Error = "internal server error"
			internal.WriteJSON(w, status, resp)
		}
		slog.Info("/update", "status", http.StatusText(status), "code", status)
		return
	}

	if req.Result != "" {
		task, err = db.UpdateResult(req.ID, req.Result)
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
				status = http.StatusNotFound

				resp.Error = "internal server error"
				internal.WriteJSON(w, status, resp)
			}
			slog.Info("/update", "status", http.StatusText(status), "code", status)
			return
		}
	}

	status = http.StatusOK
	resp = schemas.UpdateResponse{
		Ok:     true,
		ID:     task.ID,
		Status: task.Status.String(),
	}
	internal.WriteJSON(w, status, resp)

	slog.Info("/update", "status", http.StatusText(status), "code", status)
}
