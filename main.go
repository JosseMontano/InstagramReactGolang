package main
import (
	"log"
	"github.com/JosseMontano/InstagramReactGolang/handlers"
	"github.com/JosseMontano/InstagramReactGolang/bd"
)
func main(){
	if bd.CheckConnec() == 0{
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Drivers()
}