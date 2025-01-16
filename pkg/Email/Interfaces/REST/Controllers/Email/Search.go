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
	resource := resources.NewSearchMessageResource(queryParams.Get("query"), queryParams.Get("size"),
		queryParams.Get("from"), queryParams.Get("sort"))

	query, err := transform.SearchMessageResourceToQuery(*resource)
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	messages, pagination, err := c.MessageQueryService.Search(*query)
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := shared_resources.ToJSONresponse(shared_resources.NewResponsePagination(messages, pagination))
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(messages) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(response)
}
