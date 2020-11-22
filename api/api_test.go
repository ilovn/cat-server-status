package api

import (
	"cat-server-status/model"
	"cat-server-status/util"
	"encoding/json"
	"fmt"
	"testing"
)

func TestRecord(t *testing.T) {
	token := util.Token("iplc-sz")
	fmt.Println(token)

	node := model.Node{}
	node.Token = token
	node.IP = "xxx.xxx.xxx"
	node.Name = "IPLC深圳"
	node.Type = "Centos"

	bs ,_:= json.Marshal(node)
	fmt.Println(string(bs))

	//dao.InsertNode(node)
	fmt.Println(util.Token("iplc-hk"))

}
