package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

//GetList returns list of participants, including links to profiles/teams
func GetList(leagueURL string) (string, error) {
	participants, err := GetParticipants(leagueURL)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var res string
	for _, p := range participants {
		res += fmt.Sprintf(`%v. <a href="%v">%v</a> | <a href="%v">%v</a> | %v | %v | [<a href="/squad %v">SQUAD</a>]`, p.Pos, p.TeamURL, p.Team, p.ProfileURL, p.Name, p.TourPoints, p.TotalPoints, p.Team)
		res += "\n"
	}

	return res, err
}

//GetTableImg makes a snapshot of the table and returns result as an image
func GetTableImg(leagueURL string) (io.Reader, error) {
	cmd := exec.Command("phantomjs", "snapshot.js", leagueURL, "table.stat-table", "table.png")
	cmd.Run()
	f, err := os.Open("table.png")
	defer f.Close()

	return f, err
}

func GetSquad(leagueURL string, team string) (string, error){
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

	return f, err
}