package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//MongoCN is the object the conection to the BD
var MongoCN= ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jose:8021947cbba@cluster0.eqo8gqc.mongodb.net/?retryWrites=true&w=majority");

/*ConectarBD is the function to connect to the BD */
func ConectarBD() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err !=nil{
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}
//CheckConnec is the ping to the BD
func CheckConnec() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 1
}