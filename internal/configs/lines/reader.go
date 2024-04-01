package lines

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed lines.txt
var lines embed.FS

func parseLine(data []string) (*Line, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty file lines.txt")
	}

	indices := make([]int, len(data))
	for i, char := range data {
		indx, err := strconv.Atoi(char)
		if err != nil {
			return nil, fmt.Errorf("parsing error with char %s: %w", char, err)
		}
		if indx < 0 {
			return nil, fmt.Errorf("char %s must be positive indice", char)
		}

		indices[i] = indx
	}

	return NewLine(indices), nil
}

func ReadLines() (Lines, error) {
	data, err := reader.Read(lines, "lines.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	resp := make([]Line, len(data))
	for i, str := range data {
		line, err := parseLine(str)
		if err != nil {
			return nil, fmt.Errorf("parseLines(): %w", err)
		}

		resp[i] = *line
	}

	return resp, nil
}
