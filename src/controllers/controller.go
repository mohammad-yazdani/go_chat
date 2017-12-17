package controllers

import (
	"net/http"
	"io/ioutil"
)

type Controller struct {
	mapping string
	rest Rest
}

func (controller Controller) init(rest Rest) {
	controller.rest = rest
	mux := http.NewServeMux()
	mux.HandleFunc("/" + controller.mapping, controller.handle)
}
func (controller Controller) handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": controller.get(w, r)
	case "POST": controller.post(w, r)
	}
}
func (controller Controller) get(w http.ResponseWriter, r *http.Reequest) {
	if !controller.rest.getAction(w, r) {
		http.Error(w, "GET method not defined for this end-point.", http.StatusBadRequest)
	}
}

func (controller Controller) post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error getting body",
			http.StatusInternalServerError)
	}
	if !controller.rest.postAction(string(body)) {
		http.Error(w, "POST method not defined for this end-point.", http.StatusBadRequest)
	}
}