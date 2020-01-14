package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
	"github.com/willbarkoff/autoscout/config"
	"github.com/willbarkoff/autoscout/data"
	"github.com/willbarkoff/autoscout/providers/pennfirst"
)

type byExpO []data.Team

func (a byExpO) Len() int           { return len(a) }
func (a byExpO) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byExpO) Less(i, j int) bool { return a[i].ExpO > a[j].ExpO }

func main() {
	config := config.Configure()

	teams, matchesSpreadsheet := pennfirst.PennFIRST{}.GetData(config)

	// teams := map[int]Team{
	// 	4174: Team{
	// 		Number: 4174,
	// 		Name:   "Atomic Theory",
	// 	},
	// 	6051: Team{
	// 		Number: 6051,
	// 		Name:   "Quantum Mechanics",
	// 	},
	// 	9371: Team{
	// 		Number: 9371,
	// 		Name:   "General Relativity",
	// 	},
	// 	9372: Team{
	// 		Number: 9372,
	// 		Name:   "Standard Model",
	// 	},
	// 	11453: Team{
	// 		Number: 11453,
	// 		Name:   "Uncertianty Principle",
	// 	},
	// }

	// matchesSpreadsheet := [][]string{
	// 	// {"", "", "4174 6051", "11453 9371", "10", "", "", "", "", "0", "20", "", "", "", "", "0"},
	// 	// {"", "", "11453 9372", "4174 6051", "15", "", "", "", "", "0", "19", "", "", "", "", "0"},
	// 	// {"", "", "9371 4174", "6051 9372", "18", "", "", "", "", "0", "30", "", "", "", "", "0"},
	// 	// {"", "", "4174 11453", "6051 9371", "17", "", "", "", "", "0", "12", "", "", "", "", "0"},
	// }

	matches := []data.Match{}

	for _, matchRow := range matchesSpreadsheet {
		redTeamsStr := strings.Split(matchRow[2], " ")
		blueTeamsStr := strings.Split(matchRow[3], " ")
		redScore, err := strconv.Atoi(matchRow[4])
		redPenalty, err := strconv.Atoi(matchRow[9])
		blueScore, err := strconv.Atoi(matchRow[10])
		bluePenalty, err := strconv.Atoi(matchRow[15])
		if err != nil {
			panic(err)
		}

		red := redScore - redPenalty
		blue := blueScore - bluePenalty

		match := data.Match{
			Red:         red,
			Blue:        blue,
			RedScore:    redScore,
			BlueScore:   blueScore,
			RedPenalty:  redPenalty,
			BluePenalty: bluePenalty,
		}

		for _, team := range redTeamsStr {
			number, err := strconv.Atoi(team)
			if err != nil {
				panic(err)
			}
			teamStats := teams[number]
			teamStats.Scores = append(teamStats.Scores, red)
			teams[number] = teamStats
			match.RedTeam = append(match.RedTeam, number)
		}

		for _, team := range blueTeamsStr {
			number, err := strconv.Atoi(team)
			if err != nil {
				panic(err)
			}
			teamStats := teams[number]
			teamStats.Scores = append(teamStats.Scores, blue)
			teams[number] = teamStats
			match.BlueTeam = append(match.BlueTeam, number)
		}

		matches = append(matches, match)
	}

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
