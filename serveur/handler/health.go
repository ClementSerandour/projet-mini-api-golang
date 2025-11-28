package handler
import (
        "fmt"
        "net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
        jsonStatus := `{"status":"ok"}`
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, jsonStatus)
}