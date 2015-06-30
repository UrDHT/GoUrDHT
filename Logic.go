/*
# DHT Logic Module

This module is the most complex and is close to being a god object for the project.
Most refinement and modifications required for re-purposing will happen here.

The DHT Logic exposes:

- seek(id) -> return a good forward peer for an id
- getPeers() -> return my current list of 1 hop peers
- getNotified(origin) -> notify me that origin exists
- DoIOwn(id) -> returns True iff I think I am responsible for an address.

The DHT Logic Depends On:

- Network Component: Allows communication with other nodes
    - Seek(remote,id)
    - GetPeers(remote)
    - Notify(remote,origin)


DHT logic can be internally separated into two parts:
    - Reactive Logic:
        - dealing with queries
        - Has a dedicated worker, processes requests as fast as possible
    - Periodic Logic:
        - deals with maintenance
        - essentially an infinite loop with sleeps.

# TODO:
    - Decide on a good method for dealing with bad peers!!!!!

## Reactive Logic:
 Dealing with a reactive query should work as follows:
 ```
    get read lock on peerInfo
    copy needed peerInfo into local memory
    release read lock on peerInfo
    do required computation generally just a "findmax"
    return value
```

## Periodic Logic:
```
    do while running:
        get read lock on peerInfo
        make local copy
        release read lock on peerInfo
        notify all peers
        sleep for a bit
        get read lock on new_candidates
        make a local copy
        release read lock on new_candidates
        new_peerlist = pick_new_peerlist(peerInfo_copy + new_candidates_copy)
        get write lock on peerInfo
        write new_peerlist over peerInfo
        release write lock on peerInfo
        sleep for a bit

```
*/

package GoUrDHT

type DHTLogic struct {
	network        Networking
	database       DataBase
	shortPeers     []string
	longPeers      []string
	seekCandidates []string
	notifiedMe     []string
	locPeerDict    map[string]string
	info           PeerInfo
}

func (dht DHTLogic) setup(network Networking, database DataBase) bool {
	return true
}

func (dht DHTLogic) join(peers []string) bool {
	return true
}

func (dht DHTLogic) shutdown() bool {
	return true
}

func (dht DHTLogic) doIOwn(key) bool {
	return true
}

func (dht DHTLogic) seek(key) string {
	return nil
}

func (dht DHTLogic) getPeers() []string {
	return nil
}

func (dht DHTLogic) getNotified(origin string) bool {
	return true
}

func (dht DHTLogic) onResponsibilityChange() {
	return
}

func DHTJanitor(parent DHTLogic, running bool) {
	//TODO: This should be converted into a GoRoutine
}
