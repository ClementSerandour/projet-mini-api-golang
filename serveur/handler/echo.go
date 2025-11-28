package handler
//Import des packages nécessaires
import (
	"encoding/json"
    "net/http"
	"fmt"

)
//Structure pour l'endpoint ECHO
type EchoData struct {
    Value string `json:"value"`
}
//Fonction pour l'endpoint ECHO
func Echo(w http.ResponseWriter, r *http.Request){
		var i EchoData
		json.NewDecoder(r.Body).Decode(&i)
		fmt.Fprintln(w, "Value :", i.Value)
		w.Header().Set("Content-Type", "application/json")
}