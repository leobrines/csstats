package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MrWaggel/gosteamconv"
	"github.com/go-resty/resty/v2"
)

const steamUserEndpoint = "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s&format=json"

var client = resty.New()

type SteamUserResponse struct {
	Avatar  string `json:"avatarfull"`
	Profile string `json:"profileurl"`
}

func GetSteamUserData(steamid string) (*SteamUserResponse, error) {
	steamid64, err := gosteamconv.SteamStringToInt64(steamid)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf(steamUserEndpoint, os.Getenv("STEAM_API_KEY"), steamid64)

	resp, err := client.R().Get(u)
	if err != nil {
		return nil, err
	}

	r := &struct {
		Response struct {
			Players []*SteamUserResponse `json:"players"`
		} `json:"response"`
	}{}

	body := resp.Body()

	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r.Response.Players[0], err
}
