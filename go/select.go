package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Unicorn struct {
    Name  string `bson:"name"`
    //Weight int `bson:"weight"`
}

func main() {
    session, _ := mgo.Dial("mongodb://localhost")
    defer session.Close()
    // db := session.DB("db")
    db := session.DB("test")

    // name指定でユーザ取得
    var name = "Kenny"
    var results []Unicorn
    db.C("unicorns").Find(bson.M{"name": name}).All(&results)
    fmt.Println("Results of unicorn: ", results)    
}