package gamestate

import (
	"math/rand"
	"time"
)

const (
	DefaultBoardSize = 11
	DefaultStepMS    = 900
	MinStepMS        = 250
	MaxStepMS        = 1800
)

type simulatorSnake struct {
	Name   string
	Body   []Coord
	Dir    Coord
	Health int
	Alive  bool
	Score  int
	// preferred fallback turn ordering after the forward move.
	Bias []Coord
}

type Simulator struct {
	boardSize  int
	stepMS     int
	turn       int
	paused     bool
	winnerText string
	food       *Coord
	snakes     []simulatorSnake
	rng        *rand.Rand
}

func NewSimulator(seed int64) *Simulator {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	s := &Simulator{
		boardSize: DefaultBoardSize,
		stepMS:    DefaultStepMS,
		rng:       rand.New(rand.NewSource(seed)),
	}
	s.Reset()
	return s
}

func (s *Simulator) Reset() {
	s.turn = 0
	s.paused = false
	s.winnerText = ""
	s.stepMS = DefaultStepMS
	s.snakes = []simulatorSnake{
		{
			Name:   "ALPHA",
			Body:   []Coord{{X: 2, Y: 5}, {X: 1, Y: 5}, {X: 0, Y: 5}},
			Dir:    Coord{X: 1, Y: 0},
			Health: 100,
			Alive:  true,
			Score:  0,
			Bias:   []Coord{{X: 0, Y: -1}, {X: 0, Y: 1}},
		},
		{
			Name:   "BETA",
			Body:   []Coord{{X: 8, Y: 5}, {X: 9, Y: 5}, {X: 10, Y: 5}},
			Dir:    Coord{X: -1, Y: 0},
			Health: 100,
			Alive:  true,
			Score:  0,
			Bias:   []Coord{{X: 0, Y: 1}, {X: 0, Y: -1}},
		},
	}
	s.spawnFood()
}

func (s *Simulator) ApplyCommand(cmd ControlCommand) {
	switch cmd.Command {
	case "pause_toggle":
		s.paused = !s.paused
	case "reset":
		s.Reset()
	case "faster":
		s.stepMS -= 150
		if s.stepMS < MinStepMS {
			s.stepMS = MinStepMS
		}
	case "slower":
		s.stepMS += 150
		if s.stepMS > MaxStepMS {
			s.stepMS = MaxStepMS
		}
	}
}

func (s *Simulator) Step() {
	if s.paused || s.winnerText != "" {
		return
	}

	s.turn++
	occupied := map[Coord]string{}
	decisions := make([]Coord, len(s.snakes))
	nextHeads := make([]Coord, len(s.snakes))
	headCounts := map[Coord]int{}
	ateFood := false

	for _, snake := range s.snakes {
		if !snake.Alive {
			continue
		}
		for _, segment := range snake.Body {
			occupied[segment] = snake.Name
		}
	}

	for i, snake := range s.snakes {
		if !snake.Alive {
			continue
		}
		decisions[i] = s.chooseMove(snake, occupied)
		head := snake.Body[0]
		nextHeads[i] = Coord{X: head.X + decisions[i].X, Y: head.Y + decisions[i].Y}
		headCounts[nextHeads[i]]++
	}

	eliminated := map[int]bool{}
	for i, snake := range s.snakes {
		if !snake.Alive {
			continue
		}

		head := nextHeads[i]
		if headCounts[head] > 1 {
			eliminated[i] = true
			continue
		}
		if !s.onBoard(head) {
			eliminated[i] = true
			continue
		}

		occupiedBy, blocked := occupied[head]
		tail := snake.Body[len(snake.Body)-1]
		if blocked && !(occupiedBy == snake.Name && head == tail) {
			eliminated[i] = true
		}
	}

	for i := range s.snakes {
		snake := &s.snakes[i]
		if !snake.Alive {
			continue
		}
		if eliminated[i] {
			snake.Alive = false
			continue
		}

		head := nextHeads[i]
		snake.Body = append([]Coord{head}, snake.Body...)
		snake.Dir = decisions[i]
		snake.Health--

		if s.food != nil && head == *s.food {
			snake.Health = 100
			snake.Score++
			ateFood = true
		} else {
			snake.Body = snake.Body[:len(snake.Body)-1]
		}

		if snake.Health <= 0 {
			snake.Alive = false
		}
	}

	if ateFood {
		s.spawnFood()
	}
	s.finishMatchIfNeeded()
}

func (s *Simulator) Snapshot() Snapshot {
	mode := "LIVE"
	if s.paused {
		mode = "PAUSED"
	}
	if s.winnerText != "" {
		mode = "DONE"
	}

	snakes := make([]Snake, 0, len(s.snakes))
	for _, snake := range s.snakes {
		body := make([]Coord, len(snake.Body))
		copy(body, snake.Body)
		health := snake.Health
		if !snake.Alive {
			health = 0
		}
		snakes = append(snakes, Snake{
			Name:   snake.Name,
			Health: health,
			Length: len(snake.Body),
			Score:  snake.Score,
			Alive:  snake.Alive,
			Body:   body,
		})
	}

	var food *Coord
	if s.food != nil {
		f := *s.food
		food = &f
	}

	return Snapshot{
		BoardSize:  s.boardSize,
		Turn:       s.turn,
		StepMS:     s.stepMS,
		Mode:       mode,
		WinnerText: s.winnerText,
		Food:       food,
		Snakes:     snakes,
	}
}

func (s *Simulator) chooseMove(snake simulatorSnake, occupied map[Coord]string) Coord {
	currentDir := snake.Dir
	opposite := Coord{X: -currentDir.X, Y: -currentDir.Y}
	candidates := []Coord{
		currentDir,
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}

	bestMove := currentDir
	bestScore := 1 << 30
	for _, move := range candidates {
		if move == opposite {
			continue
		}
		nextPos := Coord{X: snake.Body[0].X + move.X, Y: snake.Body[0].Y + move.Y}
		score := 0
		if !s.isSafeTarget(nextPos, occupied, snake) {
			score += 10000
		}
		score += s.distanceToFood(nextPos) * 10
		score += s.biasPenalty(move, snake.Bias)
		score += s.rng.Intn(4)

		if score < bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}

func (s *Simulator) isSafeTarget(pos Coord, occupied map[Coord]string, snake simulatorSnake) bool {
	if !s.onBoard(pos) {
		return false
	}

	occupiedBy, blocked := occupied[pos]
	if !blocked {
		return true
	}

	tail := snake.Body[len(snake.Body)-1]
	return occupiedBy == snake.Name && pos == tail
}

func (s *Simulator) distanceToFood(pos Coord) int {
	if s.food == nil {
		return 0
	}
	return abs(pos.X-s.food.X) + abs(pos.Y-s.food.Y)
}

func (s *Simulator) biasPenalty(move Coord, bias []Coord) int {
	if len(bias) > 0 && move == bias[0] {
		return 0
	}
	if len(bias) > 1 && move == bias[1] {
		return 1
	}
	return 2
}

func (s *Simulator) spawnFood() {
	freeCells := make([]Coord, 0, s.boardSize*s.boardSize)
	occupied := map[Coord]bool{}
	for _, snake := range s.snakes {
		if !snake.Alive {
			continue
		}
		for _, segment := range snake.Body {
			occupied[segment] = true
		}
	}

	for y := 0; y < s.boardSize; y++ {
		for x := 0; x < s.boardSize; x++ {
			coord := Coord{X: x, Y: y}
			if !occupied[coord] {
				freeCells = append(freeCells, coord)
			}
		}
	}

	if len(freeCells) == 0 {
		s.food = nil
		return
	}

	food := freeCells[s.rng.Intn(len(freeCells))]
	s.food = &food
}

func (s *Simulator) finishMatchIfNeeded() {
	alive := make([]simulatorSnake, 0, len(s.snakes))
	for _, snake := range s.snakes {
		if snake.Alive {
			alive = append(alive, snake)
		}
	}

	if len(alive) > 1 && s.food != nil {
		return
	}

	if len(alive) == 1 {
		s.winnerText = alive[0].Name + " wins"
		return
	}
	if len(alive) == 0 {
		s.winnerText = "Draw"
		return
	}

	leader := s.leadingSnake()
	s.winnerText = leader.Name + " survives"
}

func (s *Simulator) leadingSnake() simulatorSnake {
	best := s.snakes[0]
	for _, snake := range s.snakes[1:] {
		if snake.Score > best.Score {
			best = snake
			continue
		}
		if snake.Score == best.Score && len(snake.Body) > len(best.Body) {
			best = snake
		}
	}
	return best
}

func (s *Simulator) onBoard(coord Coord) bool {
	return coord.X >= 0 && coord.X < s.boardSize && coord.Y >= 0 && coord.Y < s.boardSize
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
