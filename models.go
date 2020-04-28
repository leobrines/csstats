package main

type Player struct {
	Name         string `json:"name"`
	SteamID      string `json:"steamid"`
	Kills        int    `json:"kills"`
	Deaths       int    `json:"deaths"`
	SteamProfile string `json:"steam_profile"`
	SteamAvatar  string `json:"steam_avatar"`
}
