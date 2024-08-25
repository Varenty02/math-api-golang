package common

import "math"

type Paginated[T any] struct {
	PageIndex       int   `json:"page_index"`
	TotalPages      int   `json:"total_pages"`
	PageSize        int   `json:"page_size"`
	TotalRecord     int64 `json:"total_record"`
	Items           []T   `json:"items"`
	HasPreviousPage bool  `json:"has_previous_page"`
	HasNextPage     bool  `json:"has_next_page"`
}

// NewPaginated creates a new Paginated instance.
func NewPaginated[T any](items []T, pageIndex, pageSize, totalRecord int64) Paginated[T] {
	totalPages := int(math.Ceil(float64(totalRecord) / float64(pageSize)))
	hasPreviousPage := pageIndex > 1
	hasNextPage := pageIndex < int64(totalPages)

	return Paginated[T]{
		PageIndex:       int(pageIndex),
		TotalPages:      totalPages,
		PageSize:        int(pageSize),
		TotalRecord:     totalRecord,
		Items:           items,
		HasPreviousPage: hasPreviousPage,
		HasNextPage:     hasNextPage,
	}
}