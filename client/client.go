package client

import (
	"bytes"
	"cat-server-status/lib"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Start()  {
	send()
	t := time.NewTicker(60 * time.Second)
	for range t.C {
		go send()
	}
}
func send()  {


	var jsonStr = []byte(lib.GetState())
	fmt.Println("jsonStr", jsonStr)
	fmt.Println("new_str", bytes.NewBuffer(jsonStr))


	req, err := http.NewRequest("POST", "http://127.0.0.1:5321/record/7c488d61a704fd469d2a08bee9bbc4e0", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}