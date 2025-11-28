package handler
//Import des packages nécessaires
import (
	"runtime"
	"encoding/json"
	"net/http"
)
//Structure des métriques RAM
type ramUse struct {
    Alloc int `json:"alloc"`
	TotalAlloc int `json:"totalalloc"`
	Sys int `json:"sys"`
	NumGC int `json:"numgc"`
}
//Handler pour l'endpoint INFO
func Info(w http.ResponseWriter, r *http.Request){
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		i := ramUse{
			Alloc: int(m.Alloc),
			TotalAlloc: int(m.TotalAlloc),
			Sys: int(m.Sys),
			NumGC: int(m.NumGC),
		}
		json.NewEncoder(w).Encode(i)
	}