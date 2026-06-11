package gamestate

import (
	"math/rand"
	"time"
)

const (
	DefaultBoardSize = 11
	DefaultStepMS    = 700
	MinStepMS        = 250
	MaxStepMS        = 1800
)

type simulatorSnake struct {
	Name      string
	Archetype string
	Body      []Coord
	Dir       Coord
	Health    int
	Alive     bool
	Score     int
	// preferred fallback turn ordering after the forward move.
	Bias []Coord
	// strategy knobs
	FoodWeight    int
	EnemyWeight   int
	CenterWeight  int
	SpaceWeight   int
	RandomWeight  int
	EdgeAversion  int
	ContestWeight int
}

type Simulator struct {
	boardSize   int
	stepMS      int
	matchNumber int
	matchIndex  int
	turn        int
	paused      bool
	title       string
	eventText   string
	winnerText  string
	foods       []Coord
	snakes      []simulatorSnake
	rng         *rand.Rand
}

type snakeProfile struct {
	Name          string
	Archetype     string
	ForwardDir    Coord
	Bias          []Coord
	FoodWeight    int
	EnemyWeight   int
	CenterWeight  int
	SpaceWeight   int
	RandomWeight  int
	EdgeAversion  int
	ContestWeight int
}

var matchupProfiles = [][]snakeProfile{
	{
		{
			Name:          "FORAGER",
			Archetype:     "Greedy",
			ForwardDir:    Coord{X: 1, Y: 0},
			Bias:          []Coord{{X: 0, Y: -1}, {X: 0, Y: 1}},
			FoodWeight:    16,
			EnemyWeight:   2,
			CenterWeight:  3,
			SpaceWeight:   8,
			RandomWeight:  2,
			EdgeAversion:  4,
			ContestWeight: 7,
		},
		{
			Name:          "HUNTER",
			Archetype:     "Aggro",
			ForwardDir:    Coord{X: -1, Y: 0},
			Bias:          []Coord{{X: 0, Y: 1}, {X: 0, Y: -1}},
			FoodWeight:    7,
			EnemyWeight:   14,
			CenterWeight:  4,
			SpaceWeight:   5,
			RandomWeight:  3,
			EdgeAversion:  3,
			ContestWeight: 15,
		},
	},
	{
		{
			Name:          "LOOPER",
			Archetype:     "Safe",
			ForwardDir:    Coord{X: 1, Y: 0},
			Bias:          []Coord{{X: 0, Y: 1}, {X: 0, Y: -1}},
			FoodWeight:    8,
			EnemyWeight:   3,
			CenterWeight:  5,
			SpaceWeight:   15,
			RandomWeight:  1,
			EdgeAversion:  6,
			ContestWeight: 4,
		},
		{
			Name:          "SPRINTER",
			Archetype:     "Burst",
			ForwardDir:    Coord{X: -1, Y: 0},
			Bias:          []Coord{{X: 0, Y: -1}, {X: 0, Y: 1}},
			FoodWeight:    14,
			EnemyWeight:   6,
			CenterWeight:  2,
			SpaceWeight:   6,
			RandomWeight:  4,
			EdgeAversion:  2,
			ContestWeight: 10,
		},
	},
	{
		{
			Name:          "TRAPPER",
			Archetype:     "Cutoff",
			ForwardDir:    Coord{X: 1, Y: 0},
			Bias:          []Coord{{X: 0, Y: -1}, {X: 0, Y: 1}},
			FoodWeight:    5,
			EnemyWeight:   12,
			CenterWeight:  6,
			SpaceWeight:   10,
			RandomWeight:  2,
			EdgeAversion:  5,
			ContestWeight: 14,
		},
		{
			Name:          "SCAVENGER",
			Archetype:     "Opportunist",
			ForwardDir:    Coord{X: -1, Y: 0},
			Bias:          []Coord{{X: 0, Y: 1}, {X: 0, Y: -1}},
			FoodWeight:    11,
			EnemyWeight:   7,
			CenterWeight:  4,
			SpaceWeight:   7,
			RandomWeight:  5,
			EdgeAversion:  3,
			ContestWeight: 9,
		},
	},
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
	s.eventText = "Fresh board"
	s.winnerText = ""
	s.stepMS = DefaultStepMS
	s.matchNumber++
	matchup := matchupProfiles[s.matchIndex%len(matchupProfiles)]
	s.matchIndex++
	s.title = matchup[0].Name + " vs " + matchup[1].Name
	s.snakes = []simulatorSnake{
		snakeFromProfile(matchup[0], []Coord{{X: 2, Y: 5}, {X: 1, Y: 5}, {X: 0, Y: 5}}),
		snakeFromProfile(matchup[1], []Coord{{X: 8, Y: 5}, {X: 9, Y: 5}, {X: 10, Y: 5}}),
	}
	s.seedFoods()
}

func (s *Simulator) ApplyCommand(cmd ControlCommand) {
	switch cmd.Command {
	case "pause_toggle":
		s.paused = !s.paused
		if s.paused {
			s.eventText = "Paused"
		} else {
			s.eventText = "Back in motion"
		}
	case "reset":
		s.Reset()
	case "faster":
		s.stepMS -= 150
		if s.stepMS < MinStepMS {
			s.stepMS = MinStepMS
		}
		s.eventText = "Turn pace increased"
	case "slower":
		s.stepMS += 150
		if s.stepMS > MaxStepMS {
			s.stepMS = MaxStepMS
		}
		s.eventText = "Turn pace reduced"
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
	s.eventText = "Sizing up the next turn"

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
		decisions[i] = s.chooseMove(i, snake, occupied)
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
			s.eventText = snake.Name + " lost a head-on clash"
			continue
		}
		if !s.onBoard(head) {
			eliminated[i] = true
			s.eventText = snake.Name + " hit the wall"
			continue
		}

		occupiedBy, blocked := occupied[head]
		tail := snake.Body[len(snake.Body)-1]
		if blocked && !(occupiedBy == snake.Name && head == tail) {
			eliminated[i] = true
			if occupiedBy == snake.Name {
				s.eventText = snake.Name + " folded into itself"
			} else {
				s.eventText = snake.Name + " crashed into " + occupiedBy
			}
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

		if s.consumeFood(head) {
			snake.Health = 100
			snake.Score++
			ateFood = true
			s.eventText = snake.Name + " found food"
		} else {
			snake.Body = snake.Body[:len(snake.Body)-1]
		}

		if snake.Health <= 0 {
			snake.Alive = false
			s.eventText = snake.Name + " starved out"
		}
	}

	if ateFood {
		s.restockFoods()
	} else if len(s.foods) < 2 && s.turn%4 == 0 {
		s.spawnFood()
		s.eventText = "Bonus food popped in"
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
			Name:      snake.Name,
			Archetype: snake.Archetype,
			Health:    health,
			Length:    len(snake.Body),
			Score:     snake.Score,
			Alive:     snake.Alive,
			Body:      body,
		})
	}

	var food *Coord
	foods := make([]Coord, len(s.foods))
	copy(foods, s.foods)
	if len(foods) > 0 {
		f := foods[0]
		food = &f
	}

	return Snapshot{
		BoardSize:   s.boardSize,
		MatchNumber: s.matchNumber,
		Title:       s.title,
		EventText:   s.eventText,
		Turn:        s.turn,
		StepMS:      s.stepMS,
		Mode:        mode,
		WinnerText:  s.winnerText,
		Food:        food,
		Foods:       foods,
		Snakes:      snakes,
	}
}

func (s *Simulator) chooseMove(index int, snake simulatorSnake, occupied map[Coord]string) Coord {
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
		score += s.distanceToNearestFood(nextPos) * snake.FoodWeight
		score += s.distanceToNearestEnemyHead(index, nextPos) * snake.EnemyWeight
		score += s.distanceToCenter(nextPos) * snake.CenterWeight
		score -= s.openNeighbors(nextPos, occupied, snake) * snake.SpaceWeight
		score += s.edgePenalty(nextPos) * snake.EdgeAversion
		score += s.contestPenalty(index, nextPos) * snake.ContestWeight
		score += s.biasPenalty(move, snake.Bias)
		score += s.rng.Intn(1 + snake.RandomWeight)

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

func (s *Simulator) distanceToNearestFood(pos Coord) int {
	if len(s.foods) == 0 {
		return s.boardSize
	}
	best := 1 << 30
	for _, food := range s.foods {
		dist := abs(pos.X-food.X) + abs(pos.Y-food.Y)
		if dist < best {
			best = dist
		}
	}
	return best
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

func (s *Simulator) seedFoods() {
	s.foods = nil
	for len(s.foods) < 2 {
		s.spawnFood()
	}
}

func (s *Simulator) restockFoods() {
	for len(s.foods) < 2 {
		s.spawnFood()
	}
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
	for _, food := range s.foods {
		occupied[food] = true
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
		return
	}

	food := freeCells[s.rng.Intn(len(freeCells))]
	s.foods = append(s.foods, food)
}

func (s *Simulator) consumeFood(head Coord) bool {
	for i, food := range s.foods {
		if food == head {
			s.foods = append(s.foods[:i], s.foods[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Simulator) finishMatchIfNeeded() {
	alive := make([]simulatorSnake, 0, len(s.snakes))
	for _, snake := range s.snakes {
		if snake.Alive {
			alive = append(alive, snake)
		}
	}

	if len(alive) > 1 && len(s.foods) > 0 {
		return
	}

	if len(alive) == 1 {
		s.winnerText = alive[0].Name + " wins"
		s.eventText = alive[0].Name + " controlled the board"
		return
	}
	if len(alive) == 0 {
		s.winnerText = "Draw"
		s.eventText = "Nobody survived the exchange"
		return
	}

	leader := s.leadingSnake()
	s.winnerText = leader.Name + " survives"
	s.eventText = leader.Name + " edged the late game"
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

func (s *Simulator) openNeighbors(pos Coord, occupied map[Coord]string, snake simulatorSnake) int {
	moves := []Coord{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}
	count := 0
	for _, move := range moves {
		next := Coord{X: pos.X + move.X, Y: pos.Y + move.Y}
		if s.isSafeTarget(next, occupied, snake) {
			count++
		}
	}
	return count
}

func (s *Simulator) distanceToNearestEnemyHead(index int, pos Coord) int {
	best := 1 << 30
	for i, snake := range s.snakes {
		if i == index || !snake.Alive || len(snake.Body) == 0 {
			continue
		}
		head := snake.Body[0]
		dist := abs(pos.X-head.X) + abs(pos.Y-head.Y)
		if dist < best {
			best = dist
		}
	}
	if best == 1<<30 {
		return s.boardSize
	}
	return best
}

func (s *Simulator) distanceToCenter(pos Coord) int {
	center := s.boardSize / 2
	return abs(pos.X-center) + abs(pos.Y-center)
}

func (s *Simulator) edgePenalty(pos Coord) int {
	penalty := 0
	if pos.X == 0 || pos.X == s.boardSize-1 {
		penalty++
	}
	if pos.Y == 0 || pos.Y == s.boardSize-1 {
		penalty++
	}
	return penalty
}

func (s *Simulator) contestPenalty(index int, pos Coord) int {
	bestEnemyLength := 0
	nearest := s.boardSize * 2
	for i, snake := range s.snakes {
		if i == index || !snake.Alive || len(snake.Body) == 0 {
			continue
		}
		head := snake.Body[0]
		dist := abs(pos.X-head.X) + abs(pos.Y-head.Y)
		if dist < nearest {
			nearest = dist
			bestEnemyLength = len(snake.Body)
		}
	}
	if nearest != 1 {
		return 0
	}
	if len(s.snakes[index].Body) >= bestEnemyLength {
		return -2
	}
	return 6
}

func snakeFromProfile(profile snakeProfile, body []Coord) simulatorSnake {
	bias := make([]Coord, len(profile.Bias))
	copy(bias, profile.Bias)
	bodyCopy := make([]Coord, len(body))
	copy(bodyCopy, body)
	return simulatorSnake{
		Name:          profile.Name,
		Archetype:     profile.Archetype,
		Body:          bodyCopy,
		Dir:           profile.ForwardDir,
		Health:        100,
		Alive:         true,
		Score:         0,
		Bias:          bias,
		FoodWeight:    profile.FoodWeight,
		EnemyWeight:   profile.EnemyWeight,
		CenterWeight:  profile.CenterWeight,
		SpaceWeight:   profile.SpaceWeight,
		RandomWeight:  profile.RandomWeight,
		EdgeAversion:  profile.EdgeAversion,
		ContestWeight: profile.ContestWeight,
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
