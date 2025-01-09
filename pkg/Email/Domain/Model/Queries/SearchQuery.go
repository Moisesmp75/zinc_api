package queries

type SearchQuery struct {
	Query string
	Size  int
	From  int64
	Sort  string
}
