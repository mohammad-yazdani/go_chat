package controllers

import "net/http"

type Rest interface {
	getAction(w http.ResponseWriter, r *http.Request) bool
	postAction(data string) bool
	//put() bool
	//delete() bool
}