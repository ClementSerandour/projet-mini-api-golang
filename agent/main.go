package main
//Import des packages nécessaires
import (
	"os"
	"fmt"
	"runtime"
	"encoding/json"
	"net/http"
    "io/ioutil"
    "bytes"
    "crypto/tls"
	"time"
)
//Structure des métriques
type Metrics struct {
	Hostname string `json:"hostname"`
	OsName string `json:"osName"`
    UpTime string `json:"uptime"`
    Timestamp time.Time `json:"timestamp"`
}
//Fonction principale
func main(){
	//Heure de démarage de l'agent
	start := time.Now()
	for {
	//Récupération du Hostname
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	//Récupération de l'OS
	osName := runtime.GOOS

	//Récupération du temps d'activité
	uptime := time.Since(start)
	
	//Création de la structure Metrics
	i := Metrics{
		Hostname: host,
		OsName: osName,
		UpTime: uptime.String(),
	}

	// Encoder en JSON
    jsonData, err := json.Marshal(i)
    if err != nil {
        fmt.Println("Erreur JSON:", err)
        return
    }

	//Définition du client insecure
    client:= &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }

	//Envois en POST
	res, err := client.Post("https://localhost:3030/agent/update", "application/json", bytes.NewReader(jsonData))
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
	defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
    fmt.Println("Valeur :", string(body))
	fmt.Println(uptime)
	time.Sleep(10 * time.Second)
	}
}