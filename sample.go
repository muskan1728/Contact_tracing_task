package main
import (
   "context"
   "fmt"
   "log"
   "net/http"
   "time"
   "encoding/json"
   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/bson/primitive"
   // "go.mongodb.org/mongo-driver/mongo/readpref"
)
func ini() (context.Context,*mongo.Collection) {
   client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
   if err != nil {
      log.Fatal(err)
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
      log.Fatal(err)
  }
  defer client.Disconnect(ctx)
  trace := client.Database("Contact_tracing")
  user :=trace.Collection("Users")
  return ctx,user
}
type User struct{
   ID primitive.ObjectID
   time time.Time
   Name string
   number string
   Email string
}
func createUser(user1 *User)  {
   ctx,user :=ini()
   _,err:= user.InsertOne(ctx, user1)
   fmt.Printf("Inserted  documents into episode collection!\n")
 
   if err != nil {
      // log.Fatal(err)
      fmt.Printf("%v",err)
  }
}
func hello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/users" {
		
      fmt.Fprintf(w,"cytc")
	}
	switch r.Method {
   case "GET":		
      
       http.ServeFile(w, r, "edit.html")
       
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
      fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
      user1 := &User{
         ID : primitive.NewObjectID(),
         time : time.Now(),
         Name: r.FormValue("name"),
         number: r.FormValue("number"),
         Email:r.FormValue("email"),
     }
     createUser(user1)
     prettyJSON, err:= json.MarshalIndent(user1, "", "    ")
     fmt.Fprintf(w,"%s\n", string(prettyJSON))
     if err != nil {
      log.Fatal("Failed to generate json", err)
  }
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func search(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w,"%s",r.URL.Path[7:])
   // c.FindId(bson.M{"_id": bson.ObjectIdHex(r.URL.Path[7:])})
   ctx,user :=ini()
   var u User
   docID ,err:= primitive.ObjectIDFromHex(r.URL.Path[7:])
   fmt.Println(`bson.M{"_id": docID}:`, bson.M{"_id": docID})
   err = user.FindOne(ctx, bson.M{"_id": docID}).Decode(&u)
   if err != nil {
      // log.Fatal(err)
      err = user.FindOne(ctx, bson.M{"_id": docID}).Decode(&u)
  }
   
   // fmt.Println(`bson.M{"_id": docID}:`, bson.M{"_id": docID})

   
}
func main() {
   http.HandleFunc("/users/",search)
	http.HandleFunc("/users", hello)
   ini() 
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
