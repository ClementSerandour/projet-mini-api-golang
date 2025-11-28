//main.go
package main
import (
    "log"
    "net/http"
	"tp-mini-api/handler"
)

func main(){
	//Endpoint HEALTH
    http.HandleFunc("GET /health", handler.HealthCheck)
	
	//Endpoint TIME
	http.HandleFunc("GET /time", handler.Time)

	//Endpoint ECHO
	http.HandleFunc("POST /echo", handler.Echo)

	//Endpoint INFO
	http.HandleFunc("GET /info", handler.Info)

	//Endpoint AGENT/UPDATE
	http.HandleFunc("POST /agent/update", handler.AgentUpdate)
	
	//Endpoint AGENT/INFO
	http.HandleFunc("GET /agent/info", handler.AgentInfo)

	//starting http server
	log.Println("Starting server on :3030")
	http.ListenAndServeTLS(":3030", "certs/localhost.pem", "certs/localhost-key.pem", nil)
}