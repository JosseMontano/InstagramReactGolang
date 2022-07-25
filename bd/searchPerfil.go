package bd
import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/JosseMontano/InstagramReactGolang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func SearchPerfil(ID string) (models.User, error){
	ctx, cancel :=context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("instagram")
	col := db.Collection("users")
	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":objID,
	}
	err := col.FindOne(ctx, condition).Decode(&perfil)
	perfil.Password=""
	if err != nil{
		return perfil, err
	}
	return perfil, nil
}