package routers
import(
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/JosseMontano/InstagramReactGolang/bd"
	"github.com/JosseMontano/InstagramReactGolang/models"
)
var Email string
var IDUser string
func ProcessToken(tk string) (*models.Claim, bool, string, error){
	myKey := []byte("myKey")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2{
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return myKey, nil
	})
	if err == nil {
		 _, encontrado, _ := bd.CheckExitsUser(claims.Email)
		 if encontrado{
			Email = claims.Email
			IDUser = claims.ID.Hex()
		 }
		 return claims, encontrado, IDUser, nil
	}
	if !tkn.Valid{
		return claims, false, string(""), errors.New("token Invalido")
	}
	return claims, false, string(""), err
}