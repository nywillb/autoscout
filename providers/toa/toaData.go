package toa

type TOAMatch struct {
	MatchKey        string                `json:"match_key"`
	EventKey        string                `json:"event_key"`
	TournamentLevel int                   `json:"tournament_level"`
	ScheduledTime   string                `json:"scheduled_time"`
	MatchName       string                `json:"match_name"`
	PlayNumber      int                   `json:"play_number"`
	FieldNumber     int                   `json:"field_number"`
	PrestartTime    string                `json:"prestart_time"`
	MatchStartTime  string                `json:"match_start_time"`
	PrestartCount   int                   `json:"prestart_count"`
	CycleTime       int                   `json:"cycle_time"`
	RedScore        int                   `json:"red_score"`
	BlueScore       int                   `json:"blue_score"`
	RedPenalty      int                   `json:"red_penalty"`
	BluePenalty     int                   `json:"blue_penalty"`
	RedAutoScore    int                   `json:"red_auto_score"`
	BlueAutoScore   int                   `json:"blue_auto_score"`
	RedTeleScore    int                   `json:"red_tele_score"`
	BlueTeleScore   int                   `json:"blue_tele_score"`
	RedEndScore     int                   `json:"red_end_score"`
	BlueEndScore    int                   `json:"blue_end_score"`
	VideoURL        string                `json:"video_url"`
	Participants    []TOAMatchParticipant `json:"participants"`
}

type TOAMatchParticipant struct {
	MatchParticipantKey string  `json:"match_participant_key"`
	MatchKey            string  `json:"match_key"`
	TeamKey             string  `json:"team_key"`
	Station             int     `json:"station"`
	StationStatus       int     `json:"station_status"`
	RefStatus           int     `json:"ref_status"`
	Team                TOATeam `json:"team"`
}

type TOAEventParticipant struct {
	EventParticipantKey string  `json:"event_participant_key"`
	EventKey            string  `json:"event_key"`
	TeamKey             string  `json:"team_key"`
	TeamNumber          int     `json:"team_number"`
	IsActive            bool    `json:"is_active"`
	Team                TOATeam `json:"team"`
}

type TOATeam struct {
	TeamKey       string `json:"team_key"`
	RegionKey     string `json:"region_key"`
	LeaugeKey     string `json:"league_key"`
	TeamNumber    int    `json:"team_number"`
	TeamNameShort string `json:"team_name_short"`
	TeamNameLong  string `json:"team_name_long"`
	RobotName     string `json:"robot_name"`
	LastActive    string `json:"last_active"`
	City          string `json:"city"`
	State         string `json:"state_prov"`
	ZIPCode       string `json:"zip_code"`
	Country       string `json:"country"`
	RookieYear    int    `json:"rookie_year"`
	Website       string `json:"website"`
}
