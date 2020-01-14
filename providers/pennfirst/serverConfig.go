package pennfirst

type serverConfig struct {
	Divisions []Division
	Jsv       int
}

// Division represents a division
type Division struct {
	Name    string
	Sources DivisionSources
}

// DivisionSources represents the links to the Scoring System generated html files
type DivisionSources struct {
	Matches  string
	Rankings string
	Details  string
	Teams    string
}
