package response

import (
	"log/slog"
	"net/http"

	"github.com/harshad-salunke2002/books_api/models"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		slog.Info("responding with 5xx error", msg)
	}
	error := models.ErrorResponse{
		Error:   msg,
		Success: false,
	}

	ResponseWithJson(w, code, error)

}
