/*

This file describes the Networking class and associated helper threads/objects
see: https://github.com/UrDHT/DevelopmentPlan/blob/master/Network.md



*/

package network

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UrDHt/GoUrDHT/util"
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

// GetIP the ip of the asker, the one calling the function
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

func (nw Networking) GetPeers(remote string) []util.PeerInfo {
	var results []util.PeerInfo
	resp, err := http.Get(remote + version + peerURL + "getPeers")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	if err = json.Unmarshal(bytes, &results); err != nil {
		panic(err)
	}

	return results
}

/*
func (nw Networking) setup(logic DHTLogic, database DataBase) {
	//TODO: what is the server type?
	//server.setLinks(logic, database)
}



func (nw Networking) seek(remote, id string) string {
}



func (nw Networking) notify(remote, origin string) bool {
}

func (nw Networking) store(remote, id, data string) bool {
}

//func (nw Networking) Ping(remote string) string {

*/
