package handler
import (
	"encoding/json"
    "net/http"
	"fmt"

)

type EchoData struct {
    Value string `json:"value"`
}

func Echo(w http.ResponseWriter, r *http.Request){
		var i EchoData
		json.NewDecoder(r.Body).Decode(&i)
		fmt.Fprintln(w, "Value :", i.Value)
		w.Header().Set("Content-Type", "application/json")
}