package handlers

import (
	"errors"
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
	task, err = db.UpdateStatus(req.ID, req.NewStatus)
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

	if req.Result != "" {
		task, err = db.UpdateResult(req.ID, req.Result)
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
	}

	resp = schemas.UpdateResponse{
		Ok:     true,
		ID: task.ID,
		Status: task.Status.String(),
	}
	internal.WriteJSON(w, http.StatusOK, resp)
}
