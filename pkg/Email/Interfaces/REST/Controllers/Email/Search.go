package controllers

import (
	"log"
	resources "mamuro_api/pkg/Email/Interfaces/REST/Resources"
	transform "mamuro_api/pkg/Email/Interfaces/REST/Transform"
	shared_resources "mamuro_api/pkg/Shared/Interfaces/REST/Resources"
	"net/http"
)

func (c *MessageController) SearchMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()
	resource := resources.NewSearchMessageQuery(queryParams.Get("query"), queryParams.Get("size"),
		queryParams.Get("from"), queryParams.Get("sort"))

	query, err := transform.SearchQueryResourceToQuery(*resource)
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusBadRequest, err.Error())
		// errorResponse := shared_resources.ErrorResponse(err.Error())
		// response, _ := shared_resources.ToJSONresponse(errorResponse)
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(response)
		return
	}

	messages, pagination, err := c.MessageQueryService.Search(*query)
	if err != nil {
		log.Println(err.Error())

		shared_resources.HandleErrorResponse(w, http.StatusInternalServerError, err.Error())
		// errorResponse := shared_resources.ErrorResponse(err.Error())
		// response, _ := shared_resources.ToJSONresponse(errorResponse)
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write(response)
		return
	}

	response, err := shared_resources.ToJSONresponse(shared_resources.NewResponsePagination(messages, pagination))
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusInternalServerError, err.Error())
		// errorResponse := shared_resources.ErrorResponse(err.Error())
		// response, _ := shared_resources.ToJSONresponse(errorResponse)
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write(response)
		return
	}
	if len(messages) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(response)
}
