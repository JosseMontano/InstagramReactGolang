package middlew
import (
	"net/http"
	"github.com/JosseMontano/InstagramReactGolang/bd"
)
//CheckBD it allow know the state of the db
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if bd.CheckConnec() == 0{
			http.Error(w, "Conexion perdida con la bd", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}