// This is the REST API server.

package main

const api_version = "api/v0"

var myLogic DHTLogic
var myDB DataBase

func setLinks(logic DHTLogic, db DataBase) {
}

type RestHandler interface {
	do_HEAD()
	do_GET()
	do_POST()
}

//TODO: thread handling should be handled using
// Goroutines
