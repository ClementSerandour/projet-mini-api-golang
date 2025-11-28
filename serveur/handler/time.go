package handler
//Import des packages nécessaires
import (
	"fmt"
	"time"
    "net/http"
)
//Handler pour l'endpoint TIME
func Time(w http.ResponseWriter, r *http.Request) {
    tm := time.Now()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, tm)
	}