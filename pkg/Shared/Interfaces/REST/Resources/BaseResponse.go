package resources

import (
	"encoding/json"
	"io"
	value_objects "mamuro_api/pkg/Shared/Domain/Model/ValueObjects"
)

type BaseResponse[T any] struct {
	Success  bool
	Message  string
	Resource T
}

type BaseResponsePag[T any] struct {
	Success    bool
	Message    string
	Resource   T
	Pagination value_objects.Pagination
}

func NewResponse[T interface{}](resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Message:  "",
		Resource: resource,
	}
}

func NewResponsePagination[T interface{}](resource T, pagination value_objects.Pagination) BaseResponsePag[T] {
	return BaseResponsePag[T]{
		Success:    true,
		Message:    "",
		Resource:   resource,
		Pagination: pagination,
	}
}

func ErrorResponse(message string) BaseResponse[interface{}] {
	return BaseResponse[interface{}]{
		Success:  false,
		Message:  message,
		Resource: nil,
	}
}

func MessageResponse[T interface{}](message string, resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Message:  message,
		Resource: resource,
	}
}

func ToJSONresponse(response any) ([]byte, error) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		return []byte(""), err
	}
	return jsonData, nil
}

func DecodeRequestBody(body io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(target)
	if err != nil {
		return err
	}
	return nil
}
