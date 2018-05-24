package cmd

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/thiagotrennepohl/sdr/sdr"
)

type sdrCLI struct {
	sdrUseCase     sdr.SdrUseCase
	commaDelimiter rune
}

func NewSdrCli(sdrUseCase sdr.SdrUseCase) sdr.SdrCli {
	return &sdrCLI{
		sdrUseCase:     sdrUseCase,
		commaDelimiter: ';',
	}
}

func (s *sdrCLI) ParseHeaders(csv *csv.Reader) ([]string, error) {
	headers, err := csv.Read()
	if err != nil {
		return nil, err
	}
	headers = s.sdrUseCase.ParseHeaders(headers)
	return headers, err

}

func (s *sdrCLI) ReadCSV(filePath string) (*csv.Reader, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, FILE_NOT_FOUND_ERR
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, ERROR_OPENING_FILE
	}

	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.Comma = s.commaDelimiter

	return csvReader, err
}

func (s *sdrCLI) StoreData(csv *csv.Reader, headers []string) error {

	csvData := make([]map[string]interface{}, 0)
	for {
		line, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		data := make(map[string]interface{})
		for index, value := range line {
			data[headers[index]] = value
			csvData = append(csvData, data)
		}
	}

	s.sdrUseCase.StoreBatch
	return nil
}
