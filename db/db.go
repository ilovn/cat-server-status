package db

import (
	"log"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)
var Session *r.Session
var TS r.Term
var TB_apps r.Term
var TB_records r.Term
var TB_nodes r.Term
func init() {
	connectionPool()
	TS = r.DB("cat_server_status")
	TB_apps = TS.Table("apps")
	TB_records = TS.Table("records")
	TB_nodes = TS.Table("nodes")
}
func connectionPool() {
	var err error

	Session, err = r.Connect(r.ConnectOpts{
		Address:    "",
		InitialCap: 10,
		MaxOpen:    10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

