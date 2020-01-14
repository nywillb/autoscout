package pennfirst

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/willbarkoff/autoscout/config"
	"github.com/willbarkoff/autoscout/data"
)

// PennFIRST provides data from the format used by the Pennsylvania region
type PennFIRST struct{}

// GetData provides data
func (PennFIRST) GetData(config config.Config) (map[int]data.Team, [][]string) {
	serverConfigResp, err := http.Get(config.Stats.URL + "js/config.json")
	if err != nil {
		panic(err)
	}

	defer serverConfigResp.Body.Close()
	body, err := ioutil.ReadAll(serverConfigResp.Body)
	if err != nil {
		panic(err)
	}

	var serverConfig serverConfig

	err = json.Unmarshal(body, &serverConfig)
	if err != nil {
		panic(err)
	}

	var division Division
	for _, loopDivision := range serverConfig.Divisions {
		if loopDivision.Name == config.Stats.Division {
			division = loopDivision
		}
	}
	if (division == Division{}) {
		fmt.Println("Division " + config.Stats.Division + " not found.")
	}

	detailResp, err := http.Get(config.Stats.URL + division.Sources.Details)
	if err != nil {
		panic(err)
	}

	defer detailResp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(detailResp.Body)
	if err != nil {
		panic(err)
	}

	var matches [][]string

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		if i < 2 {
			return //Skip this iteration of the loop, because it is the header row.
		}
		row := []string{}
		s.Find("td").Each(func(j int, t *goquery.Selection) {
			row = append(row, t.Text())
		})
		matches = append(matches, row)
	})

	teamsResp, err := http.Get(config.Stats.URL + division.Sources.Teams)
	if err != nil {
		panic(err)
	}

	defer teamsResp.Body.Close()
	teamsDoc, err := goquery.NewDocumentFromReader(teamsResp.Body)
	if err != nil {
		panic(err)
	}

	var teamsArray [][]string

	teamsDoc.Find("tr").Each(func(i int, s *goquery.Selection) {
		if i < 1 {
			return //Skip this iteration of the loop, because it is the header row.
		}
		row := []string{}
		s.Find("td").Each(func(j int, t *goquery.Selection) {
			row = append(row, t.Text())
		})
		teamsArray = append(teamsArray, row)
	})

	teams := map[int]data.Team{}

	for _, team := range teamsArray {
		num, err := strconv.Atoi(team[0])
		if err != nil {
			panic(err)
		}
		thisTeam := data.Team{
			Number:      num,
			Name:        team[1],
			Affiliation: team[2],
			City:        team[3],
			State:       team[4],
			Country:     team[5],
		}
		teams[num] = thisTeam
	}

	return teams, matches
}

// GetName provides the name of the data provider
func (PennFIRST) GetName() string { return "Penn FIRST" }
