package main

type serverConfig struct {
	Divisions []Division
	Jsv       int
}

type Division struct {
	Name    string
	Sources DivisionSources
}

type DivisionSources struct {
	Matches  string
	Rankings string
	Details  string
	Teams    string
}
