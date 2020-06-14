package main

import (
	"log"
	"net/http"
)

var appname = "CaseDefinitionService"

func main() {
	StartWebServer("9090")
}

func StartWebServer(port string) {
	log.Println("Starting %v on port %v", appname, port)
	r := NewAppRouter()
	http.Handle("/" , r)
	http.ListenAndServe("localhost:" + port, r)

}
