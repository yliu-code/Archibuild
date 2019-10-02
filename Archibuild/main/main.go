package main

import (
	"net/http"
	"fmt"
	"strconv"
	"html/template"

)

func main(){
	Initialize()
}

var LATEST string = "templates/archedit.html"

var TEST_CLIENT Client = Client{&Project{TextSoFar:"Hey"},&ClientProfile{Name:"Bob",Budget:1000, Location:"Hell", Area:"4000", ArchitectKey:"Ted Mosby"}}
func search(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/search.html")
	if(err!=nil){

		fmt.Println(err);
		return
	}

	t.Execute(w,""); //change when template is generated
}



func UserView(w http.ResponseWriter, r *http.Request ){

	t, err := template.ParseFiles("templates/userView.html")
	if(err!=nil){

		fmt.Println(err);
		return
	}

	//address:=strings.Split(r.Header.Get("X-Forwarded-For"),",")



	t.Execute(w,""); //change when template is generated

}

func formFill(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/requestForm.html")

	architectNames, provided := r.URL.Query()["architect"]
	var architectName string;
	if(!provided || len(architectNames) < 1){
		architectName = "God";
	}else{
		architectName=architectNames[0]
	}


	if(err!=nil){

		fmt.Println(err);
		return
	}



	t.Execute(w,architectName); //change when template is generated
}

func index(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		notFound(w, r, http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if(err!=nil){
		fmt.Println(err);
		return
	}

	err=t.Execute(w,""); //change when template is generated

}

func bypass(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles(LATEST)
	if(err!=nil){
		fmt.Println(err);
		return
	}

	err=t.Execute(w,TEST_CLIENT); //change when template is generated
	if(err!=nil){

	}
}

func notFound(w http.ResponseWriter, r *http.Request, status int){
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Not Found")
	}
}


func clientWebServer(w http.ResponseWriter, r *http.Request){
	//private_tmpl_files := []string{"templates/client.html"}
	t, err := template.ParseFiles("templates/client.html")
	if(err!=nil){

		fmt.Println(err);
		return
	}
	t.Execute(w,""); //change when template is generated
}

func architectEdit(w http.ResponseWriter, r *http.Request){
	if(GlobalProject==nil){
		fmt.Fprintln(w,"Faliure. No Architects have a project yet");
	}else{
		http.Redirect(w,r,GlobalProject.ArchitectEditLink,http.StatusSeeOther)
	}
}

func multiplexers(handleMultiplex *http.ServeMux){
	handleMultiplex.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	handleMultiplex.HandleFunc("/author", architectEdit)
	handleMultiplex.HandleFunc("/",index)
	//handleMultiplex.HandleFunc("/aserver", StartWebserver)
	handleMultiplex.HandleFunc("/lookup", LookupServer)
	handleMultiplex.HandleFunc("/requestForm",formFill)
	handleMultiplex.HandleFunc("/sendPreferences", SendProfile)

	handleMultiplex.HandleFunc("/architectSetup",UserView)

	handleMultiplex.HandleFunc("/commun",InteractionServer)
	handleMultiplex.HandleFunc("/search", search)
	handleMultiplex.HandleFunc("/edit", bypass)
}

var Multiplex *http.ServeMux


func Initialize(){
	port:= 2323

	handleMultiplex:=http.NewServeMux()
	Multiplex = handleMultiplex
	multiplexers(handleMultiplex)

	LoadArchitects()

	server:=http.Server{Addr:"0.0.0.0:"+strconv.Itoa(port), Handler:handleMultiplex,}
	server.ListenAndServe()
}