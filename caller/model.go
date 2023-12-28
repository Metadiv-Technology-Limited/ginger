package caller

import (
	"ginger"

	"github.com/Metadiv-Technology-Limited/sql/models"
)

type Response[T any] struct {
	Success    bool               `json:"success"`
	Duration   int64              `json:"duration"`
	Pagination models.IPagination `json:"pagination,omitempty"`
	Error      ginger.IError      `json:"error,omitempty"`
	Data       *T                 `json:"data,omitempty"`
}
