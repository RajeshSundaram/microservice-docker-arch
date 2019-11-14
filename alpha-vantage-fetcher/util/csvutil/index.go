package csvutil

import (
	"bytes"
	"encoding/csv"
)

func ReadCSV(data []byte) ([][]string, error) {
	// csv.NewReader
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
