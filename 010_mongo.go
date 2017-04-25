package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const mongoHost = "localhost"
const mongoDatabase  = "test"
const mongoCollection = "objects"

type Mongo struct {
	string  //inheritance
}

type Object struct {
	Msg string
}

func main() {
	fmt.Println("start")
	c := createConn()
	insertData(c)
	result := fetchData(c)
	fmt.Println(result)
	deleteData(c)
}

func createConn() *mgo.Collection {
	session, err := mgo.Dial(mongoHost)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	fmt.Println("connection")
	return session.DB(mongoDatabase).C(mongoCollection)
	//fmt.Println(reflect.TypeOf(c))
}

func insertData(c *mgo.Collection) {
	id := 1
	msg := ""
	for id < 100 {
		msg = fmt.Sprintf("%s %d", "testowy", id)
		var err = c.Insert(&Object{msg})
		if err != nil {
			log.Fatal(err)
		}
		id = id + 1
	}
}

func fetchData(c *mgo.Collection) Object {
	var result = Object{}
	var err = c.Find(bson.M{"msg": "testowy 1"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func deleteData(c *mgo.Collection) {
	c.RemoveAll(nil)
}
