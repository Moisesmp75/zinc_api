package repositories

import (
	"fmt"
	entities "mamuro_api/pkg/Shared/Infraestructure/Persistence/Zincsearch/Entities"
)

type ZincsearchRepository[T any] struct {
	BaseUrl  string
	Username string
	Password string
	Index    string
}

func NewZincsearchRepository[T any](baseUrl string, index string, username string, password string) *ZincsearchRepository[T] {
	return &ZincsearchRepository[T]{
		BaseUrl:  baseUrl,
		Username: username,
		Password: password,
		Index:    index,
	}
}

func (r *ZincsearchRepository[T]) Insert(document T) (T, error) {

	return document, nil
}

func (r *ZincsearchRepository[T]) Search(query string, size int, from int64, sort string) (entities.SearchDataResponse[T], error) {
	url := fmt.Sprintf("%s/es/%s/_search", r.BaseUrl, r.Index)

	bodyResponse, err := performHTTPPostRequest(url, RequestToString(query, size, from, sort), r.Username, r.Password)
	if err != nil {
		return entities.SearchDataResponse[T]{}, err
	}

	var response entities.SearchDataResponse[T]

	if err := parseJSONtoDto(bodyResponse, &response); err != nil {
		return entities.SearchDataResponse[T]{}, err
	}

	return response, nil
}
