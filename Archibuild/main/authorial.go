package main

import (
	"strings"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/websocket"
	"html/template"
	"math/rand"
)

var preferenceUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Architect struct{
	Name string
	Stars int
	ImgPath string

	Projects map[string]*Project
	KeyWordTags []string
}

type Project struct{
	ArchitectEditLink string
	UserViewLink string

    TextSoFar string
    LinkToProject string

}

var UserIDs= make(map[string]*Client)
var Architects = make(map[string]*Architect)

type Client struct{
	CurrentProject *Project
	Profile *ClientProfile
}


type ClientProfile struct{
	Name string `json:"name"`
	Budget float64 `json:"budget,string"`
	Area string `json:"area"`
	Location string `json:"location"`
	Preferences string `json:"preferences"`
	ArchitectKey string`json:"architect"`
}




func LoadArchitects(){
	Architects["Ted Mosby"] = &Architect{Name:"Ted Mosby", Stars:5, Projects:make(map[string]*Project), ImgPath:"/assets/imgs/ted_mosby.jpg"}
	Architects["Bob the Builder"] = &Architect{Name:"Bob the Builder", Stars:4, Projects:make(map[string]*Project),ImgPath:"/assets/imgs/bob_builder.png"}
	Architects["Bob Parr"] = &Architect{Name:"Bob Parr", Stars:2, Projects:make(map[string]*Project),ImgPath:"/assets/imgs/Bob_Parr.jpg"}
	Architects["Jar-Jar-Binks"] = &Architect{Name:"Jar-Jar-Binks",Stars:1, Projects:make(map[string]*Project), ImgPath:"/assets/imgs/profile_jar_jar.jpg"}
}

func SearchQuery(querySearch string) string{
	querySearch = strings.ToLower(querySearch)
	resultQuery:=""
	for key := range Architects {
		lowerKey:= strings.ToLower(key)
		if(strings.HasPrefix(lowerKey,querySearch)){
			profile:=Architects[key]
			fmt.Println(profile)
			content,err :=json.Marshal(ArchQuery{Name:profile.Name, Stars:profile.Stars, Path:profile.ImgPath})
			if(err!=nil){
				break;
			}
			fmt.Println("Found Query:",string(content))
			resultQuery+=string(content)+delimter
		}
	}
	if(resultQuery==""){
		resultQuery = "null"
	}
	return resultQuery
}



/*To work on..*/
func SendProfile(w http.ResponseWriter, r *http.Request ){
	preferenceUpgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=preferenceUpgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}



	/**Prepare for new Read*/

	address:=r.Header.Get("X-Forwarded-For")
	response:=make(chan string)

	go ReadProfileUntilReceived(connection,response, address)
	go SendUntilRecieved(connection,response)


}


func ReadProfileUntilReceived(connection *websocket.Conn, response chan string, address string){
	defer connection.Close()

	for{

		_, message, err := connection.ReadMessage()
		if(err!=nil){
			fmt.Println("Failure in lookup reading")
			return
		}

		querySearch:=string(message)
		if(querySearch=="<Finished>"){
			fmt.Println("No idea")
			break;
		}

		newProfile := &ClientProfile{}
		json.Unmarshal(message,newProfile)
		fmt.Println(newProfile)

		client,ok:=UserIDs[address]
		if(!ok){
			UserIDs[address] = &Client{}
			client = UserIDs[address]
		}

		//connection.Close() // check this line later

		client.Profile = newProfile
		GenerateProject(client)

		response<-client.CurrentProject.UserViewLink

	}

}

func SendUntilRecieved(connection *websocket.Conn, response chan string){
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




func hash(s string) string {

	randomChar:=func() uint{
		return uint((90-65)*rand.Float64()) + 65
	}
	if(len(s)<5){
		s=""
		for i:=0;i<10;i++{
			s+=string(randomChar())
		}
		return s
	}


	for i:=0;i<len(s);i++{
		s= strings.Replace(s, string(s[i]),string(randomChar()), -1)
	}
	return s

}




var GlobalProject *Project = nil//this link is tested for a universal user simply a test


func GenerateProject(client *Client){
	newProject:=&Project{TextSoFar:client.Profile.Preferences}
	architect,ok:=Architects[client.Profile.ArchitectKey]
	client.CurrentProject=newProject
	if(!ok){
		newProject.UserViewLink = "/search"
		return
	}
	architect.Projects[client.Profile.Name] = newProject


	viewLinkHandler :=func(w http.ResponseWriter, r *http.Request){
		t, err := template.ParseFiles("templates/userView.html")
		if(err!=nil){
			fmt.Println(err);
			return
		}
		t.Execute(w,""); //change when template is generated
	}
	newProject.UserViewLink = "/"+hash(client.Profile.Name)

	Multiplex.HandleFunc(newProject.UserViewLink, viewLinkHandler)



	editLinkhandler :=func(w http.ResponseWriter, r *http.Request){
		t, err := template.ParseFiles(LATEST)
		if(err!=nil){
			fmt.Println(err);
			return
		}
		t.Execute(w,client); //change when template is generated
	}
	newProject.ArchitectEditLink = "/"+hash(client.Profile.Name+client.Profile.ArchitectKey)
	GlobalProject = newProject
	Multiplex.HandleFunc(newProject.ArchitectEditLink, editLinkhandler)

}
