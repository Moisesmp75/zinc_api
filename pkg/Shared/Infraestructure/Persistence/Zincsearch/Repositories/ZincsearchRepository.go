package repositories

import (
	"encoding/json"
	"fmt"
	value_objects "mamuro_api/pkg/Shared/Domain/Model/ValueObjects"
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

func ToJSON(value any) ([]byte, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func (r *ZincsearchRepository[T]) Insert(document T) (T, error) {
	url := fmt.Sprintf("%s/api/%s/_doc", r.BaseUrl, r.Index)
	documentJSON, err := ToJSON(document)
	if err != nil {
		return document, fmt.Errorf("error converting document to JSON: %w", err)
	}

	if _, err := performHTTPPostRequest(url, string(documentJSON), r.Username, r.Password); err != nil {
		return document, fmt.Errorf("error performing HTTP POST request: %w", err)
	}

	return document, nil
}

func (r *ZincsearchRepository[T]) Search(query string, size int, from int64, sort string) (entities.SearchDataResponse[T], value_objects.Pagination, error) {
	url := fmt.Sprintf("%s/es/%s/_search", r.BaseUrl, r.Index)

	bodyResponse, err := performHTTPPostRequest(url, RequestToString(query, size, from, sort), r.Username, r.Password)
	if err != nil {
		return entities.SearchDataResponse[T]{}, value_objects.Pagination{}, err
	}

	var response entities.SearchDataResponse[T]

	if err := parseJSONtoDto(bodyResponse, &response); err != nil {
		return entities.SearchDataResponse[T]{}, value_objects.Pagination{}, err
	}
	pagination := value_objects.GeneratePagination(response.Hits.Total.Value, size, from)
	return response, *pagination, nil
}
