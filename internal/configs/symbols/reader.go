package symbols

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed symbols.txt
var symbols embed.FS

const skipSymbol = -1

func parseReels(data [][]string) ([]Symbols, error) { // у нас есть тип Reels для []Symbols
	if len(data) == 0 {
		return nil, fmt.Errorf("empty file symbols.txt")
	}

	reels := make([]Symbols, len(data[0]))
	for _, line := range data {
		for i, char := range line {
			sim, err := strconv.Atoi(char)
			if err != nil {
				return nil, fmt.Errorf("parsing error with char %s: %w", char, err)
			}

			if sim != skipSymbol {
				reels[i] = append(reels[i], sim)
			}
		}
	}

	return reels, nil
}

// ReadReels - read symbols from file
func ReadReels() ([]Symbols, error) {
	// обрати внимание, что в файле symbols.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// символ -1 нужен только для выравнивания таблицы
	data, err := reader.Read(symbols, "symbols.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	symbols, err := parseReels(data)
	if err != nil {
		return nil, fmt.Errorf("parseReels(): %w", err)
	}

	return symbols, nil
}
