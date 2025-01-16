package resources

type SearchMessageResource struct {
	Query string
	Size  string
	From  string
	Sort  string
}

func NewSearchMessageResource(query string, size string, from string, sort string) *SearchMessageResource {
	return &SearchMessageResource{
		Query: query,
		Size:  size,
		From:  from,
		Sort:  sort,
	}
}
