/*
*	Date: 26th Aug 2016
*	Description: The purpose of this is to demonstrate an HTTP server serving HTML template files
*				 and static CSS/JS files
*/

package main;

import ( 
	"fmt";
	"net/http";
	"html/template"
)

//struct to use for capturing data from the client
type UserInfo struct {
	Name string;
	Age string;
}

//handler for the index file
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//path for index view
	indexView := "views/index.html";
	//index.html is parsed as a template file
	t, _ := template.ParseFiles(indexView);
	//nil because there is no data to pass to template file
	t.Execute(w, nil);
}

//handler for the /data endpoint
func dataHandler(w http.ResponseWriter, r *http.Request) {
	//path for success view
	successView := "views/success.html";
	//if the request method is not POST, reject it
	if(r.Method != "POST") {
		http.NotFound(w, r);
		return;
	}

	//parse the form data before being able to use it
	r.ParseForm();

	//Uncomment to print data to console
	//fmt.Println("name: ", r.Form["name"]);
	//fmt.Println("age: ", r.Form["age"]);

	//assign user info into a struct instance
	u := &UserInfo{Name: r.Form.Get("name"), Age: r.Form.Get("age")};

	//respond with success.html which accepts a UserInfo struct
	t, _ := template.ParseFiles(successView);
	t.Execute(w, u);
}

//main function
func main() {

	//set up static folder to serve static content
	staticFolder := http.FileServer(http.Dir("static"));

	//all static files should be requested from /static/...
	http.Handle("/static/", http.StripPrefix("/static/", staticFolder));

	//register handlers for endpoints
	http.HandleFunc("/", indexHandler);
	http.HandleFunc("/data", dataHandler);

	fmt.Println("Server is listening at 1337..\n");	

	//start the server and listen at 1337
	http.ListenAndServe(":1337", nil);

}

