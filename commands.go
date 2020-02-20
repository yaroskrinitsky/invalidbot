package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Commands struct {
	Cache CommandsCache
}

type CacheProvider interface {
	Get(key string) (string, bool)
	Add(key string, val string)
}

//GetList returns list of participants, including links to profiles/teams
func(c *Commands) GetList(leagueURL string) (string, error) {
	if val, hasValue := c.Cache.Get("GetList"+leagueURL); hasValue {
		return val, nil
	}
	participants, err := GetParticipants(leagueURL)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var res string
	for _, p := range participants {
		res += fmt.Sprintf(`%v. <a href="%v">%v</a> | <a href="%v">%v</a> | %v | %v`, p.Pos, p.TeamURL, p.Team, p.ProfileURL, p.Name, p.TourPoints, p.TotalPoints)
		res += "\n"
	}

	c.Cache.Add("GetList"+leagueURL, res)

	return res, err
}

//GetTableImg makes a snapshot of the table and returns result as an image
func(c *Commands) GetTableImg(leagueURL string) (string, error) {
	if val, hasValue := c.Cache.Get("GetTableImg" + leagueURL); hasValue {
		return val, nil
	}
	f := "table.png"
	cmd := exec.Command("phantomjs", "snapshot.js", leagueURL, "table.stat-table", f)
	cmd.Run()

	if _, err := os.Stat(f); err != nil {
		return "", err
	}

	c.Cache.Add("GetTableImg" + leagueURL, f)
	return f, nil
}

func(c *Commands) GetSquad(leagueURL string, team string) (string, error){
	if val, hasValue := c.Cache.Get("GetSquad" + leagueURL + team); hasValue {
		return val, nil
	}
	participants, err := GetParticipants(leagueURL)
	var p *Participant
	for _, pp := range participants {
		if strings.Trim(pp.Team, " ") == team {
			p = &pp
			break
		}
	}
	if p == nil {
		return "", errors.New(fmt.Sprintf("No such team in the league: %v", team))
	}

	f := p.Team + ".png"
	cmd := exec.Command("phantomjs", "snapshot.js", p.TeamURL, "div.grace.full-field", f)
	cmd.Run()
	if _, err := os.Stat(f); err != nil {
		return "", err
	}

	c.Cache.Add("GetSquad" + leagueURL + team, f)

	return f, err
}