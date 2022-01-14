package main

import (
	"io"
	"main/sharelin/data/mongo"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello world!\n")
}

func main() {
	MongoDataBase := mongo.MongoDatabase{}
	MongoDataBase.StartCheckConnection()
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

}
