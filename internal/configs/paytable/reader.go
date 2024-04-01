package paytable

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

//go:embed pay_table.txt
var payTable embed.FS

func parsePayouts(data [][]string) (map[symbols.Symbol]Payout, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty file pay_table.txt")
	}

	table := make(map[symbols.Symbol]Payout, len(data))
	for i, payouts := range data {
		for _, str := range payouts {
			payout, err := strconv.ParseUint(str, 0, 64)
			if err != nil {
				return nil, fmt.Errorf("parsing error with %s: %w", str, err)
			}

			table[i] = append(table[i], payout)
		}
	}

	return table, nil
}

func ReadPayTable() (*PayTable, error) {
	data, err := reader.Read(payTable, "pay_table.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	payouts, err := parsePayouts(data)
	if err != nil {
		return nil, fmt.Errorf("parsePayouts(): %w", err)
	}

	return NewPayTable(payouts), nil
}
