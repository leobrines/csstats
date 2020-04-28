package main

// GetTopPlayers will get stats of the top players
func GetTopPlayers() ([]*Player, error) {
	sql := "SELECT name, steamid, kills, deaths FROM csstats.csstats ORDER BY kills DESC LIMIT 10"

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var players []*Player

	for rows.Next() {
		p := &Player{}

		err := rows.Scan(&p.Name, &p.SteamID, &p.Kills, &p.Deaths)
		if err != nil {
			return nil, err
		}

		players = append(players, p)
	}

	return players, nil
}

func GetTopPlayersWithProfiles() ([]*Player, error) {
	players, err := GetTopPlayers()
	if err != nil {
		return nil, err
	}

	for _, p := range players {
		data, err := GetSteamUserData(p.SteamID)
		if err != nil {
			return nil, err
		}

		p.SteamProfile = data.Profile
		p.SteamAvatar = data.Avatar
	}

	return players, nil
}
