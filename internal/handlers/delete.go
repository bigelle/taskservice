package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

// Handles /delete endpoint
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		req    schemas.DeleteRequest
		resp   schemas.DeleteResponse
		status int
	)
	db := database.NewDB()

	err = internal.ReadJSON(r, &req)
	if err != nil {
		status = http.StatusBadRequest

		resp = schemas.DeleteResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, status, resp)

		slog.Info("/delete", "status", http.StatusText(status), "code", status)
		return
	}

	err = db.Delete(req.ID)
	if err != nil {
		resp.Ok = false

		if errors.Is(err, database.ErrInvalidData) {
			status = http.StatusBadRequest

			resp.Error = "bad request"
			internal.WriteJSON(w, status, resp)

		} else if errors.Is(err, database.ErrNoRecord) {
			status = http.StatusNotFound

			resp.Error = "not found"
			internal.WriteJSON(w, http.StatusBadRequest, resp)

		} else {
			status = http.StatusInternalServerError

			resp.Error = "internal server error"
			internal.WriteJSON(w, http.StatusNotFound, resp)
		}
		slog.Info("/delete", "status", http.StatusText(status), "code", status)
		return
	}

	status = http.StatusOK
	resp = schemas.DeleteResponse{Ok: true}
	internal.WriteJSON(w, http.StatusOK, resp)

	slog.Info("/delete", "status", http.StatusText(status), "code", status)
}
