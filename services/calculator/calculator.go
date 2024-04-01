package calculator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

// WILD - специальный символ, который может заменить любой другой символ
// он не имеет своего выигрыша, но может увеличить выигрыш за счет замены другого символа
const WILD = symbols.Symbol(0)

type Calculator struct {
	lines    lines.Lines
	payTable *paytable.PayTable
}

func NewCalculator(lines lines.Lines, payTable *paytable.PayTable) *Calculator {
	return &Calculator{lines: lines, payTable: payTable}
}

func (c *Calculator) Calculate(spinSymbols symbols.Reels) ([]win.Win, error) {
	prizes := make([]win.Win, 0)
	for _, line := range c.lines {
		indices := line.GetIndices()
		winLine := make(symbols.Symbols, 0, len(indices))

		var count int
		prev, curr := -1, -1
		for i, indx := range indices {
			curr = spinSymbols[i][indx]
			if curr != prev && curr != WILD && prev != -1 {
				break
			}

			winLine = append(winLine, curr)
			count++
			prev = curr
		}

		payout, err := c.payTable.Get(prev, count - 1)
		if err != nil {
			return nil, err
		}

		if payout > 0 {
			prize := win.NewWin(payout, winLine, prev)
			prizes = append(prizes, prize)
		}
	}

	if len(prizes) == 0 {
		return nil, nil
	}

	return prizes, nil
}
