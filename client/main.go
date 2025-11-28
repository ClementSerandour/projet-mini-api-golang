//main.go
package main
//Import des packages nécessaires
import (
    "crypto/tls"
    "fmt"
    "net/http"
    "io/ioutil"
	"io"
	"log"
    "encoding/json"
    "bytes"
)
//Structure pour l'endpoint ECHO
type Content struct {
    Value string `json:"value"`
}
//Fonction principale
func main(){
    client:= &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }

	//Interrogation endpoint HEALTH
	res, err := client.Get("https://localhost:3030/health")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	//Interrogation endpoint TIME
	res, err = client.Get("https://localhost:3030/time")
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	//Interrogation endpoint ECHO
    item := Content{Value: "Gwladys"}
    body, err = json.Marshal(item)
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        return
    }
	//Envois en POST
	res, err = client.Post("https://localhost:3030/echo", "application/json", bytes.NewReader(body))
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
    defer res.Body.Close()
    body, err = ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }
    fmt.Println("Valeur du echo :", string(body))

	//Interrogation endpoint INFO
	res, err = client.Get("https://localhost:3030/info")
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(res.Body) 
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	//Interrogation endpoint AGENT/INFO
	res, err = client.Get("https://localhost:3030/agent/info")
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
	
}