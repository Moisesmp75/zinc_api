package resources

import "encoding/json"

type BaseResponse[T any] struct {
	Success  bool
	Message  string
	Resource T
}

type BaseResponsePag[T any] struct {
	Success    bool
	Message    string
	Resource   T
	Pagination Pagination
}

func NewResponse[T interface{}](resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Message:  "",
		Resource: resource,
	}
}

func NewResponsePagination[T interface{}](resource T, pagination Pagination) BaseResponsePag[T] {
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
