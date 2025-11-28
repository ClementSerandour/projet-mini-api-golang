package handler
import (
	"runtime"
	"encoding/json"
	"net/http"
)

type ramUse struct {
    Alloc int `json:"alloc"`
	TotalAlloc int `json:"totalalloc"`
	Sys int `json:"sys"`
	NumGC int `json:"numgc"`
}

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