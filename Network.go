/*

This file describes the Networking class and associated helper threads/objects
see: https://github.com/UrDHT/DevelopmentPlan/blob/master/Network.md



*/

package main

//import "encoding/json"
import "net/http"
import "fmt"
import "io/ioutil"

var version string = "api/v0/"
var clientURL string = version+"client/"
var peerURL string = version + "peer/"

type Networking struct {
	ip, port string
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

func (nw Networking) GetIP(remote string) string {
}
*/
//func (nw Networking) Ping(remote string) string {
func Ping(remote string) {
	resp, _  := http.Get(remote+version+peerURL+"ping")
	text, _ :=  ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	
	fmt.Printf("The repsonse is %s", text)
}


func main() {
	target := "http://envy.blamestross.com:8000/"
	Ping(target)
}
