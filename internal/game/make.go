package game

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/services/calculator"
	"github.com/releaseband/golang-developer-test/services/generator"
)

const rowsCount = 3

func New() (*Slot, error) {
	winLines, err := lines.ReadLines()
	if err != nil {
		return nil, err
	}

	payTable, err := paytable.ReadPayTable()
	if err != nil {
		return nil, err
	}

	reels, err := symbols.ReadReels()
	if err != nil {
		return nil, err
	}

	gen := generator.NewSymbols(rowsCount, reels)
	calc := calculator.NewCalculator(winLines, payTable)
	slot := newSlot(gen, calc, RoundCost(rowsCount))

	return slot, nil
}
