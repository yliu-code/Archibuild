package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"net/http"
)

var lookupUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ArchQuery struct{
     Name string `json:"name"`
     Stars int   `json:"stars"`
     Path string  `json:"path"`
}


var delimter string = "</Delim>"


func LookupServer(w http.ResponseWriter, r *http.Request, ){
	lookupUpgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=lookupUpgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}

	response:=make(chan string)
	go LookUpReadUntilClose(connection, response)
	go LookupWriteUntilClose(connection,response)



}

func LookUpReadUntilClose(connection *websocket.Conn, response chan string){
	defer connection.Close()

	for{

		_, message, err := connection.ReadMessage()
		querySearch:=string(message)
		if(querySearch=="<Finished>"){
			break;
		}
		if(len(querySearch)<=1){
			response<-""
			continue;
		}

		if(err!=nil){
			fmt.Println("Failure in lookup reading")
			return
		}

		response<-SearchQuery(querySearch)
	}

}

func LookupWriteUntilClose(connection *websocket.Conn, response chan string){
	defer connection.Close()

	for{
		writer, err := connection.NextWriter(websocket.TextMessage)
		if(err!=nil){
			fmt.Println("Failure in lookup writing")
			return
		}

		writer.Write([]byte(<-response))
	}
}









