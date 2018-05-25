package sdr

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type sdr struct {
	commaDelimiter rune
}

type SdrConfig struct {
	CommaDelimiter rune
}

func NewSdr(sdrconfig SdrConfig) Sdr {
	return &sdr{
		commaDelimiter: sdrconfig.CommaDelimiter,
	}
}

func (s *sdr) ParseHeaders(csv *csv.Reader) ([]string, error) {
	headers, err := csv.Read()
	if err != nil {
		return nil, err
	}
	headers = s.fixHeaders(headers)

	return headers, err

}

func (sdr *sdr) fixHeaders(headers []string) []string {
	for index, value := range headers {
		value = strings.Trim(value, " ")
		value = strings.Replace(value, " ", "_", -1)
		headers[index] = strings.ToUpper(value)
	}
	return headers
}

func (s *sdr) ReadCSV(filePath string) (*csv.Reader, error) {
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

//StoreData
func (s *sdr) Extract(csv *csv.Reader, headers []string) ([]map[string]interface{}, error) {

	csvData := make([]map[string]interface{}, 0)
	for {
		line, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		data := make(map[string]interface{})
		for index, value := range line {
			data[headers[index]] = value
			csvData = append(csvData, data)
		}
	}

	return csvData, nil
}
