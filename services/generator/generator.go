package generator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

const gameBoardSize = 3

type Symbols struct {
	rowsCount int
	gameTapes []symbols.Symbols // symbols.Reels
}

func NewSymbols(rowsCount int, gameTapes []symbols.Symbols) *Symbols {
	return &Symbols{rowsCount: rowsCount, gameTapes: gameTapes}
}

func (s *Symbols) Generate(rng rng.RNG) (symbols.Reels, error) {
	reels := make(symbols.Reels, len(s.gameTapes))
	for i, tape := range s.gameTapes {
		indices := make(map[int]bool, len(tape))
		length := uint32(len(tape) - 1)
		for range tape {
			indx := int(rng.Random(0, length))
			if ok := indices[indx]; !ok {
				indices[indx] = true
				reels[i] = append(reels[i], tape[indx])

				continue
			}

			for shift := indx + 1; len(reels[i]) != len(tape); shift++{
				shift = shift % len(tape)
				if ok := indices[shift]; !ok {
					indices[shift] = true
					reels[i] = append(reels[i], tape[shift])

					break
				}
			}
		}
	}

	s.gameTapes = reels

	return reels, nil
}

func (s *Symbols) GetReelSymbols(reelIndex int, rowIndex int) symbols.Symbols {
	reelSymbols := make(symbols.Symbols, 0, gameBoardSize)
	for i := rowIndex; len(reelSymbols) != gameBoardSize; i++ {
		i = i % len(s.gameTapes[reelIndex])
		reelSymbols = append(reelSymbols, s.gameTapes[reelIndex][i])
	}

	return reelSymbols
}
