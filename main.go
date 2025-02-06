package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"discord-lookup/api"

	"github.com/joho/godotenv"
)

var port = 3000

func init() {
	godotenv.Load(".env.local", ".env")

	if os.Getenv("DISCORD_TOKEN") == "" {
		log.Fatal("No discord token found in .env file")
	}

	if os.Getenv("PORT") != "" {
		localport, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
		}
		port = localport
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "The id provided is not a valid discord id", 400)
			return
		}
		userInfo, err := api.GetUserInfo(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(userInfo)
	})
	router.HandleFunc("GET /v1/guild/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "The id provided is not a valid discord id", 400)
			return
		}
		guildInfo, err := api.GetGuildInfo(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(guildInfo)
	})
	router.HandleFunc("GET /v1/application/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "The id provided is not a valid discord id", 400)
			return
		}
		applicationInfo, err := api.GetApplicationInfo(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(applicationInfo)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server listening on port %d", port)
}
