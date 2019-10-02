package main

/***File Deprecated. Intended File interactionserver***/
/*
var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}




var unifyingRead = make(chan string)


//test variable-delete later
var activated bool = false;


func StartWebserver(w http.ResponseWriter, r *http.Request, ){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=upgrader.Upgrade(w,r,nil)

	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}


	connection.ReadMessage()



}

func ChatRead(connection *websocket.Conn){
	defer connection.Close()

	for{

		_, message, err := connection.ReadMessage()

		if(err!=nil){
			fmt.Println("Failure in reading")
			return
		}



		fmt.Println(string(message))
	}

}

func CharWrite(connection *websocket.Conn){
	defer connection.Close()

	for{
		writer, err := connection.NextWriter(websocket.TextMessage)
		if(err!=nil){
			fmt.Println("Failure in writing")
			return
		}

		//writer.Write([]byte(sendBack))
	}
}


*/




