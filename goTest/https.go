package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a https server\n")
	server()
}

func main() {
	http.HandleFunc("/", handler)
	// http.ListenAndServe(":8085", nil)
	http.ListenAndServeTLS(":8085", "server.crt", "server.key", nil)
}

func server() {
	client := &http.Client{}
	var data map[string]string
	data = make(map[string]string)
	data["pn"] = "13007117683"
	data["ct"] = "验证码:"
	data["secret"] = "afcfce61c1d56670d750d928a8370fc4"
	dJSON, _ := json.Marshal(data)
	resp, err := client.Post("https://cas.ipanel.cn/api/message/message_sm/send", "application/json", bytes.NewBuffer(dJSON))
	if err != nil {
		log.Printf("client.Post: %s", err.Error())
	}
	defer resp.Body.Close()
	resJSON, _ := ioutil.ReadAll(resp.Body)
	log.Printf("resJSON=%s", resJSON)
}
