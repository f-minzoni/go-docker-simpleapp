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

func main() {
  http.HandleFunc("/", list)
  http.HandleFunc("/new", add)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func list(w http.ResponseWriter, r *http.Request) {
  // get the session using information from environment, ignore errors
  session, _ := mgo.Dial(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
  db := session.DB(os.Getenv("DB_NAME"))
  defer session.Close()

  // get all records
  pings := []Ping{}
  db.C("pings").Find(nil).All(&pings)

  fmt.Fprint(w, pings)
}

func add(w http.ResponseWriter, r *http.Request) {
  // get the session using information from environment, ignore errors
  session, _ := mgo.Dial(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
  db := session.DB(os.Getenv("DB_NAME"))
  defer session.Close()

  // insert new record
  ping := Ping{
    Id:   bson.NewObjectId(),
    Time: time.Now(),
  }
  db.C("pings").Insert(ping)

  fmt.Fprint(w, ping)
}