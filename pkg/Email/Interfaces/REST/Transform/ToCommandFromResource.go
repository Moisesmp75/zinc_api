package transform

import (
	"errors"
	domain_queries "mamuro_api/pkg/Email/Domain/Model/Queries"
	resources "mamuro_api/pkg/Email/Interfaces/REST/Resources"
	"strconv"
	"strings"
)

func SearchQueryResourceToQuery(resource resources.SearchQueryResource) (*domain_queries.SearchQuery, error) {
	query := strings.TrimSpace(resource.Query)

	size := 10
	if resource.Size != "" {
		parsedSize, err := strconv.Atoi(resource.Size)
		if err != nil {
			return nil, errors.New("el parámetro 'size' debe ser un número entero")
		}
		size = parsedSize
	}

	from := int64(0)
	if resource.From != "" {
		parsedFrom, err := strconv.ParseInt(resource.From, 10, 64)
		if err != nil || parsedFrom < 0 {
			return nil, errors.New("el parámetro 'from' debe ser un número entero mayor o igual a 0")
		}
		from = parsedFrom
	}

	sort := strings.TrimSpace(resource.Sort)

	result, err := domain_queries.NewSearchQuery(query, size, from, sort)
	if err != nil {
		return nil, err
	}

	return result, nil
}
