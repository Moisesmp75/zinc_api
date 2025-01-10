package email_controller

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, err := c.MessageQueryService.Search(*query)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := shared_resources.ToJSONresponse(shared_resources.NewResponse(messages))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(messages) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(response)
}
