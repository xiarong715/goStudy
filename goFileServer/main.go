package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		fmt.Println("please input path and port")
		os.Exit(1)
	}

	port := flag.Arg(1)
	if port == "" {
		port = "88"
	}

	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	// firewall-cmd --permanent --zone=public --add-port=1314/tcp
	// ./fileserver / 1314
	// http://192.168.80.132:1314	  // 192.168.80.132     server host
	// http://localhost:1314          // on local host

	//  fileserver.exe D:\soft 1315
	//  curl -s 192.168.80.1:1315	  // 192.168.80.1		server host
}
