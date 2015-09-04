package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"
  "time"
  "log"
  "net/http"
)

type Ping struct {
  Id   bson.ObjectId `bson:"_id"`
  Time time.Time     `bson:"time"`
}

func (p Ping) String() string {
  // layouts must use the reference time Mon Jan 2 15:04:05 MST 2006
  return fmt.Sprintf("Document inserted at %v\n", p.Time.Format("3:04:05 PM"))
}

func main() {
  http.HandleFunc("/", list)
  http.HandleFunc("/new", add)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func mongoConnect() (session *mgo.Session) {
	// get the session using information from environment
  session, err := mgo.Dial(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
  
  // panics if connection error occurs
	if err != nil {
		panic(err)
	}
	
	return session
}

func list(w http.ResponseWriter, r *http.Request) {
  session := mongoConnect();
  
  // get all records
  pings := []Ping{}
  db := mongoConnect().DB(os.Getenv("DB_NAME"))
  db.C("pings").Find(nil).All(&pings)
  
  fmt.Fprint(w, pings)
  
  session.Close()
}

func add(w http.ResponseWriter, r *http.Request) {
  session := mongoConnect();
  
  // insert new record
  ping := Ping{
    Id:   bson.NewObjectId(),
    Time: time.Now(),
  }
  db := mongoConnect().DB(os.Getenv("DB_NAME"))
  db.C("pings").Insert(ping)

  fmt.Fprint(w, ping)
  
  session.Close()
}