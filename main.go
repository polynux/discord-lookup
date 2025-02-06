package main

import (
	"flag"
	"log"
	"os"

	"discord-lookup/api"

	"github.com/joho/godotenv"
	"github.com/yassinebenaid/godump"
)

func init() {
	godotenv.Load(".env.local", ".env")

	if os.Getenv("DISCORD_TOKEN") == "" {
		log.Fatal("No discord token found in .env file")
	}
}

func main() {
	var id string
	flag.StringVar(&id, "id", "", "Use local database")
	var infoFlag string
	flag.StringVar(&infoFlag, "info", "user", "Info to get: user, guild or application")
	flag.Parse()

	if id == "" || infoFlag == "" {
		log.Fatal("You need to provide both id and info flag")
	}

	var info interface{}
	var err error
	switch infoFlag {
	case "user":
		info, err = api.GetUserInfo(id)
	case "guild":
		info, err = api.GetGuildInfo(id)
	case "application":
		info, err = api.GetApplicationInfo(id)
	default:
		log.Fatal("Wrong info type")
	}

	if err != nil {
		log.Fatal(err)
	}
	godump.Dump(info)
}
