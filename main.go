package main
import (
    "fmt"
    "context"
    "time"
	"log"
	"net/http"
    "io/ioutil"
    "regexp"
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/readpref"
)
client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:8081"))
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
 err := client.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
defer cancel()
 // Ping our db connection
 err = client.Ping(context.Background(), readpref.Primary())
 if err != nil {
 log.Fatal("Couldn't connect to the database", err)
 } else {
 log.Println("Connected!")
 }
collection := client.Database("Contact_Tracing").Collection("users")
type Page struct {
    Title string
    Body  []byte
}
type User struct{
    ID primitive.ObjectID 
    time time.Time
    Name string
    number int
    Email string
}
func createUser(user *User){
    _, err=collection.InsertOne(ctx,user)
}
// func (p *Page) save() error {
//     filename := p.Title + ".txt"
//     return ioutil.WriteFile(filename, p.Body, 0600)
// }
// func loadPage(title string) (*Page,error) {
//     filename := title + ".txt"
//     body, err := ioutil.ReadFile(filename)
	
// 	if (err != nil) {
//         return nil, err
// 	}
// 	return &Page{Title: title, Body: body},nil
// }
// func viewhandle(w http.ResponseWriter, r *http.Request,title string){
// 	p, err:= loadPage(title)
//     if err != nil {
//         http.Redirect(w, r, "/edit/"+title, http.StatusFound)
//         return
//     }
//     fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }
func editHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/users" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
    }
    switch r.Method {
        case "GET":		
             http.ServeFile(w, r, "edit.html")
        case "POST":
            user := &User{
                
                ID :primitive.NewObjectID(),
                time : time.Now(),
                Name: r.FormValue("name"),
                number:r.FormValue("number"),
                Email:r.FormValue("email"),

            }
            createUser(user)
        default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")    
    // p, err := loadPage(title)
    // if err != nil {
    //     p = &Page{Title: title}
    // }
    // fmt.Fprintf(w, "<h1>Editing %s</h1>"+
    //     "<form action=\"/save/%s\" method=\"POST\">"+
    //     "<textarea name=\"body\">%s</textarea><br>"+
    //     "<input type=\"submit\" value=\"Save\">"+
    //     "</form>",
    //     p.Title, p.Title, p.Body)
}
// func saveHandler(w http.ResponseWriter, r *http.Request,]) {
    
//     body := r.FormValue("body")
//     p := &Page{Title: title, Body: []byte(body)}
// 	err := p.save()
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

// func homePage(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w,"Homepage endpoint")
// }
// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
// func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
//         m := validPath.FindStringSubmatch(r.URL.Path)
//         if m == nil {
//             http.NotFound(w, r)
//             return
//         }
//         fn(w, r, m[2])
// 	}
// }
func handleRequests(){
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    // p1.save()
    // p2, _ := loadPage("TestPage")
    // fmt.Println(string(p2.Body))
	// http.HandleFunc("/",homePage)
	http.HandleFunc("/users/", editHandler)
    // http.HandleFunc("/view/", makeHandler(viewhandle))
    // http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8081",nil))
}
func main(){
	handleRequests()

	// var x,y int=3,4
	// var z uint=uint(x)

	
}
