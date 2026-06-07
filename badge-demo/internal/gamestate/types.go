package gamestate

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Length int     `json:"length"`
	Score  int     `json:"score"`
	Alive  bool    `json:"alive"`
	Body   []Coord `json:"body"`
}

type Snapshot struct {
	BoardSize  int     `json:"board_size"`
	Turn       int     `json:"turn"`
	StepMS     int     `json:"step_ms"`
	Mode       string  `json:"mode"`
	WinnerText string  `json:"winner_text"`
	Food       *Coord  `json:"food"`
	Snakes     []Snake `json:"snakes"`
}

type ControlCommand struct {
	Seq     int64  `json:"seq"`
	Command string `json:"command"`
}
