package api

import (
	"cat-server-status/dao"
	"cat-server-status/lib"
	"cat-server-status/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func Record(w http.ResponseWriter, r *http.Request) {
	re := model.Re{}
	vars := mux.Vars(r)
	_token := vars["token"]
	token := r.Header.Get("Token")
	if len(_token) > 0 {
		token = _token
	}

	fmt.Println("Token", token)
	if len(token) <= 0 {
		//无效数据
		re.Status = "error"
		re.Msg = "无效的请求"
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusForbidden)
		w.Write(bs)
		return
	} else if !dao.CheckToken(token){
		//token无效
		re.Status = "error"
		re.Msg = "Token无效"
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusForbidden)
		w.Write(bs)
		return
	} else {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		state := lib.State{}
		json.Unmarshal(b, &state)
		state.Token = token

		dao.InsertRecord(state)

		re.Status = "ok"
		re.Msg = ""
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bs)
	}
}

// 添加节点信息
func Node(w http.ResponseWriter, r *http.Request) {
	re := model.Re{}
	vars := mux.Vars(r)
	token := vars["token"]
	// 身份认证
	_token := r.Header.Get("Token")

	if len(_token) <= 0 {
		//无效数据
		re.Status = "error"
		re.Msg = "未授权的请求"
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusForbidden)
		w.Write(bs)
		return
	}

	fmt.Println("Token", token)
	if len(token) <= 0 {
		//无效数据
		re.Status = "error"
		re.Msg = "无效的请求"
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusForbidden)
		w.Write(bs)
		return
	} else if dao.CheckToken(token){
		//token无效
		re.Status = "error"
		re.Msg = "节点已存在"
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bs)
		return
	} else {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		node := model.Node{}
		json.Unmarshal(b, &node)
		node.Token = token
		node.UserToken = _token

		dao.InsertNode(node)

		re.Status = "ok"
		re.Msg = ""
		bs, err := json.Marshal(re)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bs)
	}
}