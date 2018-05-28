package sdr

import (
	"encoding/csv"
	"os"
)

type Sdr interface {
	ReadCSV(*os.File) *csv.Reader
	ParseHeaders(*csv.Reader) ([]string, error)
	Extract(*csv.Reader, []string) ([]map[string]interface{}, error)
}
