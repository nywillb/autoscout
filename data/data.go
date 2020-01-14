package data

// Team represents a team at an event
type Team struct {
	A           float64
	Number      int
	Name        string
	Affiliation string
	City        string
	State       string
	Country     string
	Scores      []int
	Teammates   []int
	Opar        float64
	ExpO        float64
	Varience    float64
	ExpOs       []float64
	mods        []float64
}

// Match represents a match at an event
type Match struct {
	Red         int
	Blue        int
	RedScore    int
	BlueScore   int
	RedPenalty  int
	BluePenalty int
	RedTeam     []int
	BlueTeam    []int
}
