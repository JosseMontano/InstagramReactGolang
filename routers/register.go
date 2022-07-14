package routers
import (
	"encoding/json"
	"net/http"
	"github.com/JosseMontano/InstagramReactGolang/bd"
	"github.com/JosseMontano/InstagramReactGolang/models"
)
func Register(w http.ResponseWriter, r *http.Request){
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "Error en los datos recibidos" + err.Error(), 400)
		return
	}
	if len(t.Gmail) ==0 {
		http.Error(w, "Error el email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "deber ser al menos 6 caracteres", 400)
		return
	}
	_,encontrado,_ :=bd.CheckExitsUser(t.Gmail)
	if encontrado == true{
		http.Error(w, "ya existe un usuario con ese email", 400)
		return
	}
	_,status,err := bd.InserRegister(t)
	if err !=nil{
		http.Error(w, "Ocurrio un error al inentar registras" + err.Error(), 400)
		return
	}
	if status == false{
		http.Error(w, "No se logro registrar al usuario" + err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}