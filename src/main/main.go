package main

import (
    "net/http"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Phone string
}

func main() {
    http.HandleFunc("/", someFunc)
    http.HandleFunc("/sally", personController)
    http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request)  {
    w.Write([]byte("Hello universe"))
}

func personController(w http.ResponseWriter, req *http.Request) {

    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    c := session.DB("test").C("people")
    err = c.Insert(&Person{"Ale", "+494944004"},
                   &Person{"Cla", "+345433344"})
    if err != nil {
        log.Fatal(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    w.Write([]byte("Phone: " + result.Phone))
}
