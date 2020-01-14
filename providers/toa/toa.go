package toa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/willbarkoff/autoscout/config"
	"github.com/willbarkoff/autoscout/data"
)

type TOA struct{}

// GetData provides data
func (TOA) GetData(config config.Config) (map[int]data.Team, []data.Match) {
	// Create client
	client := &http.Client{}

	// Create request
	eventReq, err := http.NewRequest("GET", "https://theorangealliance.org/api/event/"+config.Stats.TOAEventKey+"/matches", nil)

	// Headers
	eventReq.Header.Add("Content-Type", "application/json")
	eventReq.Header.Add("X-TOA-Key", config.Stats.TOAKey)
	eventReq.Header.Add("X-Application-Origin", config.Stats.TOAOrigin)

	// Fetch Request
	resp, err := client.Do(eventReq)

	if err != nil {
		panic("Couldn't contact TOA: " + err.Error())
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	TOAMatches := []TOAMatch{}
	matches := []data.Match{}

	teams := map[int]data.Team{}

	if err := json.Unmarshal(respBody, &TOAMatches); err != nil {
		panic(err)
	}

	for _, TOAMatch := range TOAMatches {
		if len(TOAMatch.Participants) != 4 {
			// it's not a qualification match
			continue
		}

		match := data.Match{}
		match.BlueScore = TOAMatch.BlueScore
		match.RedScore = TOAMatch.RedScore
		match.BluePenalty = TOAMatch.BluePenalty
		match.RedPenalty = TOAMatch.RedPenalty
		match.Blue = match.BlueScore - match.BluePenalty
		match.Red = match.RedScore - match.RedPenalty
		match.BlueTeam = []int{TOAMatch.Participants[2].Team.TeamNumber, TOAMatch.Participants[3].Team.TeamNumber}
		match.RedTeam = []int{TOAMatch.Participants[0].Team.TeamNumber, TOAMatch.Participants[1].Team.TeamNumber}

		for i, participant := range TOAMatch.Participants {
			TOATeam := participant.Team
			if (!reflect.DeepEqual(teams[TOATeam.TeamNumber], data.Team{})) {
				teamStats := teams[TOATeam.TeamNumber]
				if i == 0 || i == 1 {
					teamStats.Scores = append(teamStats.Scores, match.Red)
				} else if i == 2 || i == 3 {
					teamStats.Scores = append(teamStats.Scores, match.Blue)
				}
				continue
			}

			team := data.Team{}
			team.Name = TOATeam.TeamNameShort
			team.Affiliation = TOATeam.TeamNameLong
			team.City = TOATeam.City
			team.Country = TOATeam.Country
			team.State = TOATeam.State
			team.Number = TOATeam.TeamNumber

			if i == 0 || i == 1 {
				team.Scores = []int{match.Red}
			} else if i == 2 || i == 3 {
				team.Scores = []int{match.Blue}
			}

			teams[team.Number] = team
		}

		matches = append(matches, match)
	}

	return teams, matches
}

// GetName provides the name of the data provider
func (TOA) GetName() string { return "TOA" }
