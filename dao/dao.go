package dao

import (
	"cat-server-status/db"
	"cat-server-status/lib"
	"cat-server-status/model"
	"fmt"
)

func InsertNode(node model.Node)  {
	v1, e :=db.TB_nodes.Insert(node).Run(db.Session)
	fmt.Println(v1, e)
}
func CheckToken(token string) bool  {
	t := ""
	row, _:=db.TB_nodes.Filter(map[string]interface{}{"Token": token}).Field("Token").Run(db.Session)
	defer row.Close()
	row.One(&t)
	return len(t) > 0
}

func InsertRecord(state lib.State)  {
	v1, e :=db.TB_records.Insert(state).Run(db.Session)
	fmt.Println(v1, e)
}
