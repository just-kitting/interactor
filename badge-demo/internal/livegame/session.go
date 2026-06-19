package livegame

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"badgesnake/internal/gamestate"
	"badgesnake/internal/player"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/BattlesnakeOfficial/rules/client"
	"github.com/BattlesnakeOfficial/rules/maps"
)

const transportErrorCause = "transport-error"

type Config struct {
	Width               int
	Height              int
	Seed                int64
	TimeoutMS           int
	StepMS              int
	LiveName            string
	OpponentName        string
	OpponentMoveMode    string
	LiveBus             string
	LiveAddress         int
	ProbeInterval       time.Duration
	FoodSpawnChance     int
	MinimumFood         int
	HazardDamagePerTurn int
	ShrinkEveryNTurns   int
	GameType            string
	MapName             string
}

type Session struct {
	cfg         Config
	gameMap     maps.GameMap
	ruleset     rules.Ruleset
	matchNumber int
	gameID      string
	stepMS      int
	mode        string
	title       string
	eventText   string
	winnerText  string
	paused      bool
	nextProbeAt time.Time
	boardState  *rules.BoardState
	slots       []snakeSlot
}

type snakeSlot struct {
	ID         string
	Name       string
	Archetype  string
	Player     player.Player
	Metadata   client.SnakeMetadataResponse
	LastMove   string
	LastError  string
	LiveTarget bool
}

func NewSession(cfg Config, livePlayer player.Player, opponentPlayer player.Player) (*Session, error) {
	if livePlayer == nil {
		return nil, fmt.Errorf("live player is required")
	}
	if opponentPlayer == nil {
		return nil, fmt.Errorf("opponent player is required")
	}

	cfg = withDefaults(cfg)

	gameMap, err := maps.GetMap(cfg.MapName)
	if err != nil {
		return nil, fmt.Errorf("load map %q: %w", cfg.MapName, err)
	}

	settings := map[string]string{
		rules.ParamFoodSpawnChance:     strconv.Itoa(cfg.FoodSpawnChance),
		rules.ParamMinimumFood:         strconv.Itoa(cfg.MinimumFood),
		rules.ParamHazardDamagePerTurn: strconv.Itoa(cfg.HazardDamagePerTurn),
		rules.ParamShrinkEveryNTurns:   strconv.Itoa(cfg.ShrinkEveryNTurns),
	}

	session := &Session{
		cfg:     cfg,
		gameMap: gameMap,
		ruleset: rules.NewRulesetBuilder().
			WithSeed(cfg.Seed).
			WithParams(settings).
			WithSolo(false).
			NamedRuleset(cfg.GameType),
		stepMS: cfg.StepMS,
		slots: []snakeSlot{
			{
				ID:         "zepto-live",
				Name:       cfg.LiveName,
				Archetype:  "I2C",
				Player:     livePlayer,
				LastMove:   "up",
				LiveTarget: true,
			},
			{
				ID:        "sim-opponent",
				Name:      cfg.OpponentName,
				Archetype: titleCase(cfg.OpponentMoveMode),
				Player:    opponentPlayer,
				LastMove:  "up",
			},
		},
	}
	session.Reset()
	return session, nil
}

func (s *Session) Reset() {
	s.gameID = ""
	s.paused = false
	s.mode = "WAITING"
	s.eventText = s.waitingText("")
	s.winnerText = ""
	s.boardState = nil
	s.nextProbeAt = time.Time{}
	for i := range s.slots {
		s.slots[i].LastMove = "up"
		s.slots[i].LastError = ""
	}
	s.refreshTitle()
}

func (s *Session) ApplyCommand(cmd gamestate.ControlCommand) {
	switch cmd.Command {
	case "pause_toggle":
		switch s.mode {
		case "LIVE":
			s.paused = true
			s.mode = "PAUSED"
			s.eventText = "Paused"
		case "PAUSED":
			s.paused = false
			s.mode = "LIVE"
			s.eventText = "Live game"
		}
	case "reset":
		s.Reset()
	case "faster":
		if s.stepMS > gamestate.MinStepMS {
			s.stepMS -= 150
			if s.stepMS < gamestate.MinStepMS {
				s.stepMS = gamestate.MinStepMS
			}
		}
	case "slower":
		if s.stepMS < gamestate.MaxStepMS {
			s.stepMS += 150
			if s.stepMS > gamestate.MaxStepMS {
				s.stepMS = gamestate.MaxStepMS
			}
		}
	}
}

func (s *Session) EnsureStarted(ctx context.Context, now time.Time) error {
	if s.boardState != nil || s.mode == "LIVE" || s.mode == "PAUSED" || s.mode == "DONE" {
		return nil
	}
	if !s.nextProbeAt.IsZero() && now.Before(s.nextProbeAt) {
		return nil
	}
	s.nextProbeAt = now.Add(s.cfg.ProbeInterval)

	if err := s.probeMetadata(ctx); err != nil {
		s.mode = "WAITING"
		s.eventText = s.waitingText(err.Error())
		return nil
	}

	state, err := maps.SetupBoard(s.gameMap.ID(), s.ruleset.Settings(), s.cfg.Width, s.cfg.Height, s.snakeIDs())
	if err != nil {
		return err
	}
	gameOver, state, err := s.ruleset.Execute(state, nil)
	if err != nil {
		return err
	}
	if gameOver {
		return fmt.Errorf("unexpected game-over during initialization")
	}

	s.refreshTitle()
	s.gameID = fmt.Sprintf("badge-snake-%d-%d", s.matchNumber+1, time.Now().UnixNano())
	s.boardState = state

	for _, slot := range s.slots {
		request := s.requestForSlot(slot)
		if _, err := slot.Player.Start(ctx, request); err != nil {
			s.boardState = nil
			s.mode = "WAITING"
			s.eventText = s.waitingText(fmt.Sprintf("%s /start failed: %v", slot.Name, err))
			return nil
		}
	}

	s.matchNumber++
	s.mode = "LIVE"
	s.eventText = "Live game"
	s.winnerText = ""
	return nil
}

func (s *Session) Step(ctx context.Context) error {
	if s.boardState == nil || s.mode != "LIVE" || s.paused {
		return nil
	}

	boardState, err := maps.PreUpdateBoard(s.gameMap, s.boardState, s.ruleset.Settings())
	if err != nil {
		return err
	}

	moves := make([]rules.SnakeMove, 0, len(s.slots))
	turnFailures := make([]string, 0, len(s.slots))
	for _, slot := range s.slots {
		snake := findSnake(boardState, slot.ID)
		if snake == nil || snake.EliminatedCause != rules.NotEliminated {
			continue
		}

		move, moveErr := slot.Player.Move(ctx, s.requestForSlotWithState(*slotRef(s, slot.ID), boardState))
		if moveErr != nil {
			rules.EliminateSnake(snake, transportErrorCause, "", boardState.Turn+1)
			s.updateSlotError(slot.ID, moveErr.Error())
			turnFailures = append(turnFailures, fmt.Sprintf("%s failed", slot.Name))
			continue
		}

		move.Move = normalizeMove(move.Move)
		s.updateSlotMove(slot.ID, move.Move)
		moves = append(moves, rules.SnakeMove{ID: slot.ID, Move: move.Move})
	}

	gameOver, nextState, err := s.ruleset.Execute(boardState, moves)
	if err != nil {
		return err
	}

	nextState, err = maps.PostUpdateBoard(s.gameMap, nextState, s.ruleset.Settings())
	if err != nil {
		return err
	}
	nextState.Turn++
	s.boardState = nextState

	if len(turnFailures) > 0 {
		s.eventText = strings.Join(turnFailures, ", ")
	} else {
		s.eventText = "Live game"
	}

	if gameOver {
		s.finishGame(ctx)
	}
	return nil
}

func (s *Session) Snapshot() gamestate.Snapshot {
	s.refreshTitle()

	snakes := make([]gamestate.Snake, 0, len(s.slots))
	for _, slot := range s.slots {
		snakes = append(snakes, s.snapshotSnake(slot))
	}

	foods := make([]gamestate.Coord, 0)
	var firstFood *gamestate.Coord
	turn := 0
	if s.boardState != nil {
		turn = s.boardState.Turn
		for _, food := range s.boardState.Food {
			coord := gamestate.Coord{X: food.X, Y: food.Y}
			if firstFood == nil {
				copyCoord := coord
				firstFood = &copyCoord
			}
			foods = append(foods, coord)
		}
	}

	return gamestate.Snapshot{
		BoardSize:   s.cfg.Width,
		MatchNumber: s.matchNumber,
		Title:       s.title,
		EventText:   s.eventText,
		Turn:        turn,
		StepMS:      s.stepMS,
		Mode:        s.mode,
		WinnerText:  s.winnerText,
		Food:        firstFood,
		Foods:       foods,
		Snakes:      snakes,
	}
}

func (s *Session) probeMetadata(ctx context.Context) error {
	for i := range s.slots {
		metadata, err := s.slots[i].Player.Metadata(ctx)
		if err != nil {
			s.slots[i].LastError = err.Error()
			if s.slots[i].LiveTarget {
				return err
			}
			continue
		}
		s.slots[i].Metadata = metadata
		s.slots[i].LastError = ""
	}
	return nil
}

func (s *Session) requestForSlot(slot snakeSlot) client.SnakeRequest {
	return s.requestForSlotWithState(slot, s.boardState)
}

func (s *Session) requestForSlotWithState(slot snakeSlot, boardState *rules.BoardState) client.SnakeRequest {
	var youSnake rules.Snake
	if boardState != nil {
		for _, snake := range boardState.Snakes {
			if snake.ID == slot.ID {
				youSnake = snake
				break
			}
		}
	}

	return client.SnakeRequest{
		Game: client.Game{
			ID:      s.gameID,
			Timeout: s.cfg.TimeoutMS,
			Ruleset: client.Ruleset{
				Name:     s.ruleset.Name(),
				Version:  "badgesnake",
				Settings: client.ConvertRulesetSettings(s.ruleset.Settings()),
			},
			Map: s.gameMap.ID(),
		},
		Turn:  boardTurn(boardState),
		Board: convertStateToBoard(boardState, s.slots),
		You:   convertRulesSnake(youSnake, slot),
	}
}

func (s *Session) finishGame(ctx context.Context) {
	s.mode = "DONE"
	s.paused = false

	if s.boardState != nil {
		for _, slot := range s.slots {
			_, _ = slot.Player.End(ctx, s.requestForSlot(slot))
		}
	}

	winners := make([]string, 0, len(s.slots))
	for _, snake := range s.boardState.Snakes {
		if snake.EliminatedCause == rules.NotEliminated {
			winners = append(winners, slotRef(s, snake.ID).Name)
		}
	}
	switch len(winners) {
	case 0:
		s.winnerText = "Draw"
	case 1:
		s.winnerText = winners[0] + " wins"
	default:
		s.winnerText = strings.Join(winners, ", ") + " survive"
	}
	s.eventText = ""
}

func (s *Session) snapshotSnake(slot snakeSlot) gamestate.Snake {
	snapshot := gamestate.Snake{
		Name:      slot.Name,
		Archetype: slot.Archetype,
		Alive:     false,
		Body:      []gamestate.Coord{},
	}

	if s.boardState == nil {
		return snapshot
	}

	snake := findSnake(s.boardState, slot.ID)
	if snake == nil {
		return snapshot
	}

	snapshot.Health = snake.Health
	snapshot.Length = len(snake.Body)
	snapshot.Score = len(snake.Body)
	snapshot.Alive = snake.EliminatedCause == rules.NotEliminated
	for _, segment := range snake.Body {
		snapshot.Body = append(snapshot.Body, gamestate.Coord{X: segment.X, Y: segment.Y})
	}
	return snapshot
}

func (s *Session) snakeIDs() []string {
	ids := make([]string, 0, len(s.slots))
	for _, slot := range s.slots {
		ids = append(ids, slot.ID)
	}
	return ids
}

func (s *Session) refreshTitle() {
	s.title = s.slots[0].Name + " vs " + s.slots[1].Name
}

func (s *Session) waitingText(reason string) string {
	target := fmt.Sprintf("Waiting for %s on i2c-%s@0x%02x", s.cfg.LiveName, s.cfg.LiveBus, s.cfg.LiveAddress)
	if strings.TrimSpace(reason) == "" {
		return target
	}
	return target + ": " + trimForBanner(reason)
}

func (s *Session) updateSlotError(id string, errText string) {
	for i := range s.slots {
		if s.slots[i].ID == id {
			s.slots[i].LastError = errText
			return
		}
	}
}

func (s *Session) updateSlotMove(id string, move string) {
	for i := range s.slots {
		if s.slots[i].ID == id {
			s.slots[i].LastMove = move
			s.slots[i].LastError = ""
			return
		}
	}
}

func findSnake(boardState *rules.BoardState, id string) *rules.Snake {
	if boardState == nil {
		return nil
	}
	for i := range boardState.Snakes {
		if boardState.Snakes[i].ID == id {
			return &boardState.Snakes[i]
		}
	}
	return nil
}

func slotRef(s *Session, id string) *snakeSlot {
	for i := range s.slots {
		if s.slots[i].ID == id {
			return &s.slots[i]
		}
	}
	return &snakeSlot{}
}

func boardTurn(boardState *rules.BoardState) int {
	if boardState == nil {
		return 0
	}
	return boardState.Turn
}

func convertRulesSnake(snake rules.Snake, slot snakeSlot) client.Snake {
	body := client.CoordFromPointArray(snake.Body)
	head := client.Coord{}
	if len(snake.Body) > 0 {
		head = client.CoordFromPoint(snake.Body[0])
	}

	return client.Snake{
		ID:      snake.ID,
		Name:    slot.Name,
		Health:  snake.Health,
		Body:    body,
		Head:    head,
		Length:  len(snake.Body),
		Shout:   "",
		Latency: "0",
		Customizations: client.Customizations{
			Color: fallback(slot.Metadata.Color, "#888888"),
			Head:  fallback(slot.Metadata.Head, "default"),
			Tail:  fallback(slot.Metadata.Tail, "default"),
		},
	}
}

func convertStateToBoard(boardState *rules.BoardState, slots []snakeSlot) client.Board {
	if boardState == nil {
		return client.Board{}
	}

	activeSnakes := make([]client.Snake, 0, len(boardState.Snakes))
	for _, snake := range boardState.Snakes {
		if snake.EliminatedCause != rules.NotEliminated {
			continue
		}
		for _, slot := range slots {
			if slot.ID == snake.ID {
				activeSnakes = append(activeSnakes, convertRulesSnake(snake, slot))
				break
			}
		}
	}

	return client.Board{
		Height:  boardState.Height,
		Width:   boardState.Width,
		Food:    client.CoordFromPointArray(boardState.Food),
		Hazards: client.CoordFromPointArray(boardState.Hazards),
		Snakes:  activeSnakes,
	}
}

func normalizeMove(move string) string {
	switch strings.ToLower(strings.TrimSpace(move)) {
	case "up", "down", "left", "right":
		return strings.ToLower(strings.TrimSpace(move))
	default:
		return "up"
	}
}

func trimForBanner(text string) string {
	text = strings.TrimSpace(text)
	if len(text) <= 72 {
		return text
	}
	return text[:69] + "..."
}

func titleCase(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	value = strings.ToLower(value)
	return strings.ToUpper(value[:1]) + value[1:]
}

func fallback(value string, alt string) string {
	if strings.TrimSpace(value) == "" {
		return alt
	}
	return value
}

func withDefaults(cfg Config) Config {
	if cfg.Width == 0 {
		cfg.Width = gamestate.DefaultBoardSize
	}
	if cfg.Height == 0 {
		cfg.Height = gamestate.DefaultBoardSize
	}
	if cfg.TimeoutMS == 0 {
		cfg.TimeoutMS = 500
	}
	if cfg.StepMS == 0 {
		cfg.StepMS = gamestate.DefaultStepMS
	}
	if strings.TrimSpace(cfg.LiveName) == "" {
		cfg.LiveName = "Zepto"
	}
	if strings.TrimSpace(cfg.OpponentName) == "" {
		cfg.OpponentName = "Clocky"
	}
	if strings.TrimSpace(cfg.OpponentMoveMode) == "" {
		cfg.OpponentMoveMode = "clockwise"
	}
	if cfg.ProbeInterval <= 0 {
		cfg.ProbeInterval = 1 * time.Second
	}
	if cfg.FoodSpawnChance == 0 {
		cfg.FoodSpawnChance = 15
	}
	if cfg.MinimumFood == 0 {
		cfg.MinimumFood = 1
	}
	if cfg.HazardDamagePerTurn == 0 {
		cfg.HazardDamagePerTurn = 14
	}
	if cfg.ShrinkEveryNTurns == 0 {
		cfg.ShrinkEveryNTurns = 25
	}
	if strings.TrimSpace(cfg.GameType) == "" {
		cfg.GameType = "standard"
	}
	if strings.TrimSpace(cfg.MapName) == "" {
		cfg.MapName = "standard"
	}
	return cfg
}
