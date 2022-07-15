package jwt

import (
	"time"

	"github.com/JosseMontano/InstagramReactGolang/models"
	jwt "github.com/dgrijalva/jwt-go"
)
func GeneroJWT(t models.User) (string, error){
	myKey := []byte("NombredelByteToken")
	payload:=jwt.MapClaims{
		"gmail":t.Gmail,
		"name": t.Name,
		"last_name": t.LastName,
		"date_birth":t.DateBirth,
		"bibliograpy":t.Bibliography,
		"ubication":t.Ubication,
		"siteweb":t.SiteWeb,
		"_id":t.ID.Hex(),
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil{	
		return tokenStr, err
	}
	return tokenStr, nil

}