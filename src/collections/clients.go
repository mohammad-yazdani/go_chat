package collections

import (
	b64 "encoding/base64"
)

var clients map[string] string

func push(name string) bool {
	_, exists := clients[name]
	if exists {
		return false
	} else {
		encoded := b64.StdEncoding.EncodeToString([]byte(name))
		clients[name] = encoded
		return true
	}
}

func remove(name string) bool {
	_, exists := clients[name]
	if exists {
		delete(clients, name)
		return true
	} else {
		return false
	}
}

func get(name string) (string, bool)  {
	encoded, exits := clients[name]
	return encoded, exits
}