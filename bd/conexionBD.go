package bd

import (
  "context"
  "log"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)
// TODO change db pwd
var pwd = "pwd"
var clientOptions = options.Client().ApplyURI("mongodb+srv://dgmz:"+ pwd +"@cluster0.yucbq.mongodb.net/tuitr?retryWrites=true&w=majority")

var MongoCN = ConectarBD();

/* ConectarBD creates a client to connect to database */
func ConectarBD () *mongo.Client {
  client, err := mongo.Connect(context.TODO(), clientOptions)
  if (err != nil) {
    log.Fatal(err.Error())
    return client
  }
  err = client.Ping(context.TODO(), nil)
  if (err != nil) {
    log.Fatal(err.Error())
    return client
  }
  log.Println("Conexión exitosa a la BD")
  return client
}

/* CheckConnection pings to database */
func CheckConnection() int {
  err := MongoCN.Ping(context.TODO(), nil)
  if (err != nil) {
    return 0
  }
  return 1
}