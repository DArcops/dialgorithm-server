package v1

import (
	"net/http"

	"github.com/darcops/dialgorithm-server/models"
)

var (
	Err = map[error]int{
		models.ErrNotFound:   http.StatusNotFound,
		models.ErrToCreate:   http.StatusInternalServerError,
		models.ErrDuplicated: http.StatusConflict,
	}
)
