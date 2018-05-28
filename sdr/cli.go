package sdr

import (
	"encoding/csv"
)

type Sdr interface {
	ReadCSV(string) (*csv.Reader, error)
	ParseHeaders(*csv.Reader) ([]string, error)
	Extract(*csv.Reader, []string) ([]map[string]interface{}, error)
}
