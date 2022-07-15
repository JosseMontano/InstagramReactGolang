package routers
import (
	"encoding/json"
	"net/http"
	"github.com/JosseMontano/InstagramReactGolang/models"
	"github.com/JosseMontano/InstagramReactGolang/bd"
	"github.com/JosseMontano/InstagramReactGolang/jwt"
)
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")
	var t models.User
	err:=json.NewDecoder(r.Body).Decode(&t)
	if err != nil{	
		http.Error(w, "Usuario y/o contraseña invalidos " + err.Error(), 400)
		return
	}
	if len(t.Gmail) == 0{
		http.Error(w, "el email es requerido", 400)
		return
	}
	doc, exits :=bd.IntentLogin(t.Gmail, t.Password)
	if exits == false{
		http.Error(w, "Usuario y/o contraseña invalidos " + err.Error(), 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(doc)
	if err != nil{
		http.Error(w, "Ocurrio un error al generar el token " + err.Error(), 400)
		return
	}
	resp := models.RespLogin {
		Token : jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*expirationTime:= time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:"token",
		Value: jwtKey,
		Expires: expirationTime,
	})*/

}