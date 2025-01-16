package controllers

import (
	"log"
	resources "mamuro_api/pkg/Email/Interfaces/REST/Resources"
	transform "mamuro_api/pkg/Email/Interfaces/REST/Transform"
	shared_resources "mamuro_api/pkg/Shared/Interfaces/REST/Resources"
	"net/http"
)

func (c *MessageController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resource resources.CreateMessageResource
	if err := shared_resources.DecodeRequestBody(r.Body, &resource); err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	command := transform.CreateMessageResourceToCommand(resource)

	result, err := c.MessageCommandService.Post(*command)
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := shared_resources.ToJSONresponse(shared_resources.NewResponse(result))
	if err != nil {
		log.Println(err.Error())
		shared_resources.HandleErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
