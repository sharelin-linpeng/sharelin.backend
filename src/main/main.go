package main

import (
	"log"
	"main/config"
	_ "main/sharelin/imgbed"
	"net/http"
)

const serverPort = "server.port"
const serverHost = "server.host"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	listenAddress := config.CONFIG.Get(serverHost) + ":" + config.CONFIG.Get(serverPort)
	log.Printf("server listen address %s\n", listenAddress)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
