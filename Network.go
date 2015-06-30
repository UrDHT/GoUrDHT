/*

This file describes the Networking class and associated helper threads/objects
see: https://github.com/UrDHT/DevelopmentPlan/blob/master/Network.md



*/

package GoUrDHT

import "encoding/json"

type Networking struct {
	ip, port string
}

func (nw Networking) setup(logic DHTLogic, database DataBase) {
	//TODO: what is the server type?
	server.setLinks(logic, database)
}

func (nw Networking) seek(remote, id string) string {
}

func (nw Networking) getPeers(remote string) []string {
}

func (nw Networking) notify(remote, origin string) bool {
}

func (nw Networking) store(remote, id, data string) bool {
}

func (nw Networking) getIP(remote string) string {
}

func (nw Networking) ping(remote string) string {
}
