package handlers

import (
	"net/http"

	"github.com/bigelle/taskservice/internal"
	"github.com/bigelle/taskservice/internal/database"
	"github.com/bigelle/taskservice/internal/schemas"
)

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	var req schemas.DeleteRequest
	var resp schemas.DeleteResponse
	db := database.NewDB()

	err = internal.ReadJSON(r, &req)
	if err != nil {
		resp = schemas.DeleteResponse{
			Ok:    false,
			Error: "bad request",
		}
		internal.WriteJSON(w, http.StatusBadRequest, resp)
		return
	}

	err = db.Delete(req.TaskID)
	if err != nil {
		resp = schemas.DeleteResponse{
			Ok:    false,
			Error: "internal server error",
		}
		internal.WriteJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp = schemas.DeleteResponse{Ok: true}
	internal.WriteJSON(w, http.StatusOK, resp)
}
