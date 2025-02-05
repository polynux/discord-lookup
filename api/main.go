package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Info int

const (
	Guild Info = iota
	User
	Application
)

func getDiscordJson(i Info, id string) (interface{}, error) {
	var url string
	log.Printf("%d", i)
	switch i {
	case Guild:
		url = "https://canary.discord.com/api/v10/guilds/%s/widget.json"
	case User:
		url = "https://canary.discord.com/api/v10/users/%s"
	case Application:
		url = "https://canary.discord.com/api/v10/applications/%s/rpc"
	}

	url = fmt.Sprintf(url, id)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bot "+os.Getenv("DISCORD_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	htmldata, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var value interface{}
	err = json.Unmarshal(htmldata, &value)

	return value, err
}

func GetGuildInfo(id string) (*interface{}, error) {
	info, err := getDiscordJson(Guild, id)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetUserInfo(id string) (*interface{}, error) {
	info, err := getDiscordJson(User, id)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetApplicationInfo(id string) (*interface{}, error) {
	info, err := getDiscordJson(Application, id)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
