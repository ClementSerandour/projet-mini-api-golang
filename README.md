# TP Projet Mini API – Agent / Serveur / Client

## Description générale

Ce projet met en place un mini-système distribué en Go, comprenant trois composants :

1. **Agent** : Collecte des métriques locales de la machine (hostname, OS, uptime) et les envoie périodiquement au serveur central via HTTP POST.
2. **Serveur** : Centralise les métriques des agents, expose des endpoints HTTP sécurisés (HTTPS avec certificats auto-signés) et calcule un statut UP/WARNING/DOWN pour chaque agent.
3. **Client** : Interroge les endpoints du serveur et affiche les données collectées.

Le système utilise JSON pour l’échange de données et TLS pour sécuriser les communications.

---

## Arborescence des fichiers

```
.
├── agent
│   └── main.go        # Agent qui collecte et envoie les métriques
├── serveur
│   ├── main.go        # Serveur principal exposant les endpoints
│   └── handler
│       ├── agent.go   # Handlers pour AGENT/UPDATE et AGENT/INFO
│       ├── echo.go    # Handler pour l’endpoint ECHO
│       ├── health.go  # Handler pour l’endpoint HEALTH
│       ├── info.go    # Handler pour l’endpoint INFO
│       └── time.go    # Handler pour l’endpoint TIME
├── certs
│   ├── localhost.pem
│   └── localhost-key.pem
└── client
    └── main.go        # Client pour tester et afficher les endpoints
```

---

## Fichiers principaux

### 1. `agent/main.go`

* Collecte hostname, OS et uptime.
* Encode les métriques en JSON.
* Envoie les données toutes les 10 secondes au serveur via `POST /agent/update`.
* Gère TLS avec `InsecureSkipVerify: true` pour le développement local.

### 2. `serveur/main.go`

* Expose les endpoints :

  * `GET /health` : Vérifie que le serveur est actif.
  * `GET /time` : Retourne l’heure actuelle du serveur.
  * `POST /echo` : Retourne la valeur envoyée en JSON.
  * `GET /info` : Informations générales du serveur.
  * `POST /agent/update` : Reçoit les métriques d’un agent.
  * `GET /agent/info` : Retourne les métriques des agents avec leur statut calculé.
* Utilise TLS avec les certificats présents dans `certs/`.

### 3. `handler/agent.go`

* Stocke les métriques des agents dans une map globale `Agents`.
* Ajoute un timestamp à chaque métrique pour calculer le statut.
* Statuts :

  * **UP** : Dernière métrique reçue il y a moins de 30 secondes.
  * **WARNING** : Dernière métrique reçue entre 30 et 60 secondes.
  * **DOWN** : Plus de 60 secondes sans nouvelles métriques.
* `AgentInfo` formate les données en JSON lisible (`MarshalIndent`).

### 4. `client/main.go`

* Interroge tous les endpoints du serveur (HEALTH, TIME, ECHO, INFO, AGENT/INFO) via HTTPS.
* Affiche les réponses dans la console.
* Gère TLS avec `InsecureSkipVerify: true` pour développement local.
* Exemple d’utilisation de `POST /echo` avec envoi de données JSON.

---

## Instructions pour exécution

### Serveur

```bash
cd serveur
go run main.go
```

* Serveur disponible sur `https://localhost:3030/`.

### Agent

```bash
cd agent
go run main.go
```

* Envoie automatiquement les métriques toutes les 10 secondes au serveur.

### Client

```bash
cd client
go run main.go
```

* Affiche les données collectées sur le serveur et teste les endpoints.

---

## Notes

* Les certificats TLS sont auto-signés pour le développement local.
* Les communications sont en HTTPS, mais le client ignore la vérification du certificat (`InsecureSkipVerify: true`).
* Le projet est conçu pour un apprentissage pratique de Go, des handlers HTTP, JSON et TLS.
* Pour simplifier la logique des statuts, seule la dernière métrique reçue par agent est conservée.

---

## Améliorations possibles

* Historique complet des métriques dans un slice ou une base de données.
* Endpoints gRPC pour échanges plus efficaces.
* Commandes depuis le client pour modifier la fréquence d’envoi de l’agent.
* Interface utilisateur graphique ou CLI améliorée pour visualiser l’état des agents.