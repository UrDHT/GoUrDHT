/* This file describes the Database struct and associated helper threads/objects
see: https://github.com/UrDHT/DevelopmentPlan/blob/master/Database.md
*/

package GoUrDHT

import "time"

type DataBase struct {
	records map[string]string
	streams map[string][]string
}

// Golang doesn't have native tuples, so let's invent
// our own
type Tuple struct {
	id, val interface{}
}

func (db Database) setup() {
	return nil
}

func (db Database) shutdown() {
	return nil
}

func (db DataBase) get(id string) string {
	if val, exists := db.records[id]; exists {
		return val
	}
	return nil
}

func (db DataBase) store(id string, val string) {
	db.records[id] = val
}

func (db DataBase) post(id string, val string) {
	pair := Tuple{time.Unix(), val}
	if _, exists := db.streams[id]; exists {
		db.streams[id] = append(db.streams[id], pair)
	} else {
		db.streams[id] = pair
	}
}

func (db DataBase) poll(id string, t Time) []Tuple {
	var result []Tuple

	for key, value := range db {
		if value > t {
			pair := Tuple{key, value}
			result = append(result, pair)
		}
	}
	return result
}

func (db DataBase) getRecords() []string {
	var keys []string
	for k, _ := range db {
		keys = append(keys, k)
	}
	return keys
}
