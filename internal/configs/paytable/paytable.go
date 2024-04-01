package paytable

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

// Payout - таблица выплаты определенного символа
type Payout []uint64

// PayTable - таблица выплат всех символов
type PayTable struct {
	symbolPayouts map[symbols.Symbol]Payout
}

func NewPayTable(symbolPayouts map[symbols.Symbol]Payout) *PayTable {
	return &PayTable{symbolPayouts: symbolPayouts}
}

func (p *PayTable) Get(s symbols.Symbol, index int) (uint64, error) {
	var payout uint64
	if payouts, ok := p.symbolPayouts[s]; ok {
		payout = payouts[index]
	}

	return payout, nil
}
