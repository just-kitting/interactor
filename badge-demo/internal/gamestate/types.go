package gamestate

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	Name      string  `json:"name"`
	Archetype string  `json:"archetype"`
	Health    int     `json:"health"`
	Length    int     `json:"length"`
	Score     int     `json:"score"`
	Alive     bool    `json:"alive"`
	Body      []Coord `json:"body"`
}

type Snapshot struct {
	BoardSize   int     `json:"board_size"`
	MatchNumber int     `json:"match_number"`
	Title       string  `json:"title"`
	EventText   string  `json:"event_text"`
	Turn        int     `json:"turn"`
	StepMS      int     `json:"step_ms"`
	Mode        string  `json:"mode"`
	WinnerText  string  `json:"winner_text"`
	Food        *Coord  `json:"food"`
	Foods       []Coord `json:"foods"`
	Snakes      []Snake `json:"snakes"`
}

type ControlCommand struct {
	Seq     int64  `json:"seq"`
	Command string `json:"command"`
}
