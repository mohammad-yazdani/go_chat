package controllers

import "net/http"

type messageRest struct {}
func (messageRest messageRest) getAction(w http.ResponseWriter,
	r *http.Request) bool { return false }

func (messageRest messageRest) postAction(data string) bool {
	// TODO : Send message from guest

	return true
}

var _ = Controller{mapping: "message", rest: messageRest{}}