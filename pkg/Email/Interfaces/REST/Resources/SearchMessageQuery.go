package resources

type SearchQueryResource struct {
	Query string
	Size  string
	From  string
	Sort  string
}

func NewSearchMessageQuery(query string, size string, from string, sort string) *SearchQueryResource {
	return &SearchQueryResource{
		Query: query,
		Size:  size,
		From:  from,
		Sort:  sort,
	}
}
