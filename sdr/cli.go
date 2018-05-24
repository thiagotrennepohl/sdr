package sdr

import (
	"encoding/csv"
)

type SdrCli interface {
	ReadCSV(string) (*csv.Reader, error)
	ParseHeaders(*csv.Reader) ([]string, error)
	StoreData(*csv.Reader, []string) error
}
