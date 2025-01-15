package resources

import "net/http"

func HandleErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	errorResponse := ErrorResponse(errorMessage)
	response, _ := ToJSONresponse(errorResponse)
	w.WriteHeader(statusCode)
	w.Write(response)
}
