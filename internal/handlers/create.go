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

	if r.Method != http.MethodPost {
		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "method not allowed",
		}
		internal.WriteJSON(w, http.StatusMethodNotAllowed, resp)
		return
	}

	var req schemas.CreateRequest
	err = internal.ReadJSON(r, &req)
	if err != nil {
		slog.Error(err.Error())

		resp = schemas.CreateResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, http.StatusMethodNotAllowed, resp)
		return
	}

	var taskID uint
	taskID, err = db.Create(
		req.Name,
		req.Desciption,
	)
	if err != nil {
		slog.Error(err.Error())

		resp = schemas.CreateResponse{
			Ok: false,
		}

		if errors.Is(err, database.ErrInvalidData) {
			resp.Error = "bad request"
			internal.WriteJSON(w, http.StatusBadRequest, resp)
			return
		} else {
			resp.Error = "internal server error"
			internal.WriteJSON(w, http.StatusInternalServerError, resp)
			return
		}
	}

	resp = schemas.CreateResponse{
		Ok:   true,
		Name: req.Name,
		ID:   taskID,
	}
	internal.WriteJSON(w, http.StatusOK, resp)
	slog.Info("/create", "status", http.StatusOK)
}
