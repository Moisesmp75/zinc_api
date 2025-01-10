package resources

type Pagination struct {
	TotalItems   int64
	ItemsPerPage int
	CurrentPage  int64
	TotalPages   int64
	HasNextPage  bool
	HasPrevPage  bool
}

func GenerateMeta(totalItems int64, size int, from int64) *Pagination {
	totalPages := (totalItems + int64(size) - 1) / int64(size)
	currentPage := from/int64(size) + 1
	hasNextPage := currentPage < totalPages
	hasPrevPage := currentPage > 1
	return &Pagination{
		TotalItems:   totalItems,
		ItemsPerPage: size,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		HasNextPage:  hasNextPage,
		HasPrevPage:  hasPrevPage,
	}
}
