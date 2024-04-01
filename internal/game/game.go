package game

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/result"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

// RoundCost - функция, которая возвращет стоимость одного раунда
func RoundCost(linesCount int) uint64 {
	return uint64(linesCount)
}

type Slot struct {
	generator  Generator
	calculator Calculator
	roundCost  uint64
}

func newSlot(generator Generator, calculator Calculator, roundCost uint64) *Slot {
	return &Slot{generator: generator, calculator: calculator, roundCost: roundCost}
}

func (s *Slot) Spin(rng rng.RNG) (*result.Round, error) {
	reels, err := s.generator.Generate(rng)
	if err != nil {
		return nil, err
	}

	gameBoard := make(symbols.Reels, len(reels))
	for i, reel := range reels {
		indx := int(rng.Random(0, uint32(len(reel) - 1)))
		spinSymbols := s.generator.GetReelSymbols(i, indx)
		gameBoard[i] = spinSymbols
	}

	prizes, err := s.calculator.Calculate(gameBoard)
	if err != nil {
		return nil, err
	}

	round := result.NewRound(gameBoard, prizes, s.roundCost)

	return round, nil
}

func (s *Slot) RoundCost() uint64 {
	return s.roundCost
}
