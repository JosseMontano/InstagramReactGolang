package bd
import (
	"context"
	"time"
	"github.com/JosseMontano/InstagramReactGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)
func CheckExitsUser(gmail string) (models.User, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()
	db := MongoCN.Database("instagram")
	col := db.Collection("users")
	condicion := bson.M{"gmail":gmail}
	var res models.User
	err := col.FindOne(ctx, condicion).Decode(&res)
	ID := res.ID.Hex()
	if err != nil{
		return res, false, ID
	}
	return res, true, ID
}