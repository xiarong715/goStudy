package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ip := flag.String("ip", "0.0.0.0", "IP地址")
	port := flag.String("port", "1316", "端口")
	flag.Parse()

	http.HandleFunc("/", handle)
	fmt.Printf("bind %s:%s\n", *ip, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *ip, *port), nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
