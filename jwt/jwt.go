package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/JosseMontano/InstagramReactGolang/models"

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
		"siteWeb":t.SiteWeb,
		"_id":t.ID.Hex(),
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil{	
		return tokenStr, err
	}
	return tokenStr, nil

}