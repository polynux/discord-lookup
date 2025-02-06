package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Info int

type JsonObj map[string]interface{}

const (
	Guild Info = iota
	User
	Application
)

func getDiscordJson(i Info, id string) (map[string]interface{}, error) {
	var url string
	log.Printf("%d", i)
	switch i {
	case Guild:
		url = "https://discord.com/api/v10/guilds/%s/widget.json"
	case User:
		url = "https://discord.com/api/v10/users/%s"
	case Application:
		url = "https://discord.com/api/v10/applications/%s/rpc"
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
	var value map[string]interface{}
	err = json.Unmarshal(htmldata, &value)

	return value, err
}

type GuildInfo struct {
	Id            string
	Name          string
	InstantInvite string
	PresenceCount float64
}

func GetGuildInfo(id string) (*GuildInfo, error) {
	info, err := getDiscordJson(Guild, id)
	if err != nil {
		return nil, err
	}

	if info["code"] != nil {
		return nil, errors.New("The guild is either non-existent, unavailable, or has Server Widget/Discovery disabled.")
	}

	guildInfo := GuildInfo{
		Id:            id,
		Name:          info["name"].(string),
		InstantInvite: info["instant_invite"].(string),
		PresenceCount: info["presence_count"].(float64),
	}

	return &guildInfo, nil
}

type Avatar struct {
	Id         string
	Link       string
	IsAnimated bool
}

type Banner struct {
	Id         string
	Link       string
	IsAnimated bool
	Color      string
}

type UserInfo struct {
	Id     string
	Tag    string
	Badges []string
	Avatar Avatar
	Banner Banner
}

func GetUserInfo(id string) (*UserInfo, error) {
	info, err := getDiscordJson(User, id)
	if err != nil {
		return nil, err
	}

	var avatarLink string
	if info["avatar"] != nil {
		avatarLink = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s", id, info["avatar"])
	}

	var banner Banner
	if info["banner"] != nil {
		var bannerLink string
		if info["banner"] != nil {
			bannerLink = fmt.Sprintf("https://cdn.discordapp.com/banners/%s/%s", id, info["banner"])
		}
		banner = Banner{
			Id:         info["banner"].(string),
			Link:       bannerLink,
			IsAnimated: info["banner"] != nil && strings.HasPrefix(info["banner"].(string), "a_"),
			Color:      info["banner_color"].(string),
		}
	}

	userInfo := UserInfo{
		Id:  id,
		Tag: info["username"].(string),
		Avatar: Avatar{
			Id:         info["avatar"].(string),
			Link:       avatarLink,
			IsAnimated: info["avatar"] != nil && strings.HasPrefix(info["avatar"].(string), "a_"),
		},
		Banner: banner,
		Badges: nil,
	}

	return &userInfo, nil
}

func GetApplicationInfo(id string) (*map[string]interface{}, error) {
	info, err := getDiscordJson(Application, id)
	if err != nil {
		return nil, err
	}

	if info["code"] != nil {
		return nil, errors.New("The id provided is probably not an application.")
	}

	return &info, nil
}
