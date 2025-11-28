package handler
import (
    "net/http"
	"encoding/json"
    "time"
)

type Metrics struct {
	Hostname string `json:"hostname"`
	OsName string `json:"osName"`
    UpTime string `json:"uptime"`
    Timestamp time.Time `json:"timestamp"`
}

var (
    Agents = make(map[string]Metrics)
)
//Fonction pour la récupération des métriques envoyer par l'agent et stockage de la dernière dans le map Agents
func AgentUpdate(w http.ResponseWriter, r *http.Request){
    var m Metrics
    err := json.NewDecoder(r.Body).Decode(&m)
    if err != nil {
       http.Error(w, "Bad request", http.StatusBadRequest)
       return
    }
    m.Timestamp = time.Now()
    //Stocker les métriques dans la map Agents
	Agents[m.Hostname] =  m

    // Retourner les données reçues en JSON
    json.NewEncoder(w).Encode(m)
}
//Fonction d'affichage des métriques avec le status de l'agents, UP-WARNING-DOWN
func AgentInfo(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")

    type MetricsWithStatus struct {
        Hostname string `json:"hostname"`
        OsName   string `json:"osName"`
        UpTime   string `json:"uptime"`
        Status   string `json:"status"`
    }
    //Ajout des métriques et du status dans la map résult
    result := []MetricsWithStatus{}
    for _, m := range Agents {
        result = append(result, MetricsWithStatus{
            Hostname: m.Hostname,
            OsName:   m.OsName,
            UpTime:   m.UpTime,
            Status:   GetStatus(m),
        })
    }
    //Modification de l'indentation du résultat de sortie - aider par l'ia
    jsonData, _ := json.MarshalIndent(result, "", "  ")
    w.Write(jsonData)
}
/* Fonction de calcul du status de l'agent :
    UP si communication <10s, 
    Warning si pas de nouvelle depuis 30s,D
    Down si +30s sans réponse */
func GetStatus(agent Metrics) string {
    elapsed := time.Since(agent.Timestamp)
    if elapsed < 10*time.Second {
        return "UP"
    } else if elapsed < 30*time.Second {
        return "WARNING"
    } else {
        return "DOWN"
    }
}