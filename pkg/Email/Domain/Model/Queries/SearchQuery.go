package queries

import (
	"errors"
	"strings"
)

type SearchQuery struct {
	Query string
	Size  int
	From  int64
	Sort  string
}

func NewSearchQuery(query string, size int, from int64, sort string) (*SearchQuery, error) {
	if strings.ContainsAny(query, `'"<>`) {
		return nil, errors.New("query contiene caracteres no permitidos")
	}
	if size <= 0 || size%5 != 0 {
		return nil, errors.New("size debe ser un número entero positivo y múltiplo de 5")
	}
	if from < 0 {
		return nil, errors.New("from no puede ser menor a 0")
	}
	if strings.ContainsAny(sort, `'"<>`) {
		return nil, errors.New("query contiene caracteres no permitidos")
	}
	return &SearchQuery{
		Query: query,
		Size:  size,
		From:  from,
		Sort:  sort,
	}, nil
}
