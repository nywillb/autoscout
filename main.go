package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/montanaflynn/stats"
	"github.com/willbarkoff/autoscout/config"
	"github.com/willbarkoff/autoscout/data"
	"github.com/willbarkoff/autoscout/providers"
	"github.com/willbarkoff/autoscout/providers/pennfirst"
	"github.com/willbarkoff/autoscout/providers/toa"
)

type byExpO []data.Team

func (a byExpO) Len() int           { return len(a) }
func (a byExpO) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byExpO) Less(i, j int) bool { return a[i].ExpO > a[j].ExpO }

func main() {
	config := config.Configure()

	sources := []providers.Provider{pennfirst.PennFIRST{}, toa.TOA{}}

	var source providers.Provider

	for _, src := range sources {
		if src.GetName() == config.Stats.Type {
			source = src
			break
		}
	}

	if source == nil {
		panic("Provider not found")
	}

	teams, matches := source.GetData(config)

	for i, team := range teams {
		a := float64(0)
		for _, score := range team.Scores {
			a += float64(score)
		}
		a = a / float64(len(team.Scores))
		updatedTeam := team
		updatedTeam.A = a
		teams[i] = updatedTeam
	}

	for _, match := range matches {
		for i, team := range match.RedTeam {
			otherTeam := 0
			if i == 0 {
				otherTeam = 1
			}
			mod := teams[team].A - teams[match.RedTeam[otherTeam]].A
			expo := teams[team].A/2 + mod
			teamObj := teams[team]
			teamObj.ExpOs = append(teamObj.ExpOs, expo)
			teams[team] = teamObj
		}
		for i, team := range match.BlueTeam {
			otherTeam := 0
			if i == 0 {
				otherTeam = 1
			}
			mod := teams[team].A - teams[match.BlueTeam[otherTeam]].A
			expo := teams[team].A/2 + mod
			teamObj := teams[team]
			teamObj.ExpOs = append(teamObj.ExpOs, expo)
			teams[team] = teamObj
		}
	}

	avgBot := float64(0)

	for i, team := range teams {
		team.ExpO = avg(team.ExpOs)
		avgBot += team.ExpO
		statsData := stats.LoadRawData(team.ExpOs)
		team.Varience, _ = stats.Variance(statsData)
		// if err != nil {
		// 	panic(err)
		// }
		teams[i] = team
	}

	avgBot /= float64(len(teams))

	teamsArr := []data.Team{}

	for _, team := range teams {
		team.Opar = team.ExpO / avgBot
		teamsArr = append(teamsArr, team)
	}

	sort.Sort(byExpO(teamsArr))

	file, err := os.Create("scout.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	row := []string{
		"Number",
		"Name",
		"Affiliation",
		"City",
		"State",
		"Country",
		"ExpO",
		"Opar",
		"Var",
	}
	if err := csvWriter.Write(row); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, team := range teamsArr {
		row := []string{
			strconv.Itoa(team.Number),
			team.Name,
			team.Affiliation,
			team.City,
			team.State,
			team.Country,
			strconv.FormatFloat(team.ExpO, 'f', 2, 32),
			strconv.FormatFloat(team.Opar, 'f', 2, 32),
			strconv.FormatFloat(team.Varience, 'f', 2, 32),
		}
		err := csvWriter.Write(row)
		if err != nil {
			panic(err)
		}
	}
}

func avg(numbers []float64) float64 {
	total := float64(0)
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}
