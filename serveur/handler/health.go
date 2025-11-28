package handler
//Import des packages nécessaires
import (
        "fmt"
        "net/http"
)
//Handler pour l'endpoint HEALTH
func HealthCheck(w http.ResponseWriter, r *http.Request) {
        jsonStatus := `{"status":"ok"}`
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, jsonStatus)
}