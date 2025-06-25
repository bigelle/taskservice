package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	db := database.NewDB()
	var resp schemas.CreateResponse
	var status int

	if r.Method != http.MethodPost {
		status = http.StatusMethodNotAllowed

		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "method not allowed",
		}
		internal.WriteJSON(w, status, resp)

		slog.Info("/create", "status", http.StatusText(status), "code", status)
		return
	}

	var req schemas.CreateRequest
	err = internal.ReadJSON(r, &req)
	if err != nil {
		status = http.StatusBadRequest

		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, status, resp)

		slog.Info("/create", "status", http.StatusText(status), "code", status)
		return
	}

	var taskID uint
	taskID, err = db.Create(
		req.Name,
		req.Desciption,
	)
	if err != nil {

		resp = schemas.CreateResponse{
			Ok: false,
		}

		if errors.Is(err, database.ErrInvalidData) {
			status = http.StatusBadRequest

			resp.Error = "bad request"
			internal.WriteJSON(w, status, resp)
		} else {
			status = http.StatusInternalServerError

			resp.Error = "internal server error"
			internal.WriteJSON(w, status, resp)
		}

		slog.Info("/create", "status", http.StatusText(status), "code", status)
	}

	status = http.StatusOK
	resp = schemas.CreateResponse{
		Ok:   true,
		Name: req.Name,
		ID:   taskID,
	}
	internal.WriteJSON(w, http.StatusOK, resp)

	slog.Info("/create", "status", http.StatusText(status), "code", status)
}
