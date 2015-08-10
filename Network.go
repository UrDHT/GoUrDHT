/*

This file describes the Networking class and associated helper threads/objects
see: https://github.com/UrDHT/DevelopmentPlan/blob/master/Network.md



*/

package main

//import "encoding/json"
import (
	"fmt"
	"net/http"
)

import "io/ioutil"

var version = "api/v0/"
var clientURL = version + "client/"
var peerURL = version + "peer/"

// Networking is a struct reprenting the object responsible for out netowrking functions
type Networking struct {
	ip, port string
}

// Ping checks for the existance of the peer at remote address
// If it  fails, all other functions will too
func (nw Networking) Ping(remote string) string {
	resp, err := http.Get(remote + version + peerURL + "ping")
	if err != nil {
		return "BAD"
	}
	text, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "BAD"
	}

	return string(text)
}

func (nw Networking) GetIP(remote string) string {
	resp, err := http.Get(remote + version + peerURL + "getmyIP")
	if err != nil {
		return "BAD"
	}
	text, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "BAD"
	}

	return string(text)
}

/*
func (nw Networking) setup(logic DHTLogic, database DataBase) {
	//TODO: what is the server type?
	//server.setLinks(logic, database)
}

func (nw Networking) seek(remote, id string) string {
}

func (nw Networking) getPeers(remote string) []string {
}

func (nw Networking) notify(remote, origin string) bool {
}

func (nw Networking) store(remote, id, data string) bool {
}

//func (nw Networking) Ping(remote string) string {

*/

func main() {
	target := "http://envy.blamestross.com:8000/"
	nw := new(Networking)
	reply := nw.GetIP(target)
	fmt.Println(reply)
}
