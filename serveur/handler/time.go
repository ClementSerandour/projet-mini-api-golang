package handler
import (
	"fmt"
	"time"
    "net/http"
)

func Time(w http.ResponseWriter, r *http.Request) {
    tm := time.Now()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, tm)
	}