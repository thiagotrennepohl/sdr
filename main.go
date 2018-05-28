package main

import (
	"errors"
	"log"
	"os"

	"github.com/thiagotrennepohl/sdr/sdr"
	mgo "gopkg.in/mgo.v2"
)

var mongoAddress = os.Getenv("MONGO_ADDR")
var csvDelimiter = os.Getenv("CSV_DELIMITER")
var filePath = os.Getenv("FILE_PATH")

func init() {
	if mongoAddress == "" {
		mongoAddress = "mongodb://localhost:27017/yawoen"
	}
	if len(csvDelimiter) > 1 {
		panic(errors.New("Csv delimiter should be only one char e.g: ;"))
	}
	if csvDelimiter == "" {
		csvDelimiter = ";"
	}
	if filePath == "" {
		panic(errors.New("File path variable is empty"))
	}
}

func main() {
	session, _ := mgo.Dial(mongoAddress)

	conn := session.DB("").C("companies")
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: csvDelimiter})

	csv, err := sdr.ReadCSV(filePath)
	haders, err := sdr.ParseHeaders(csv)
	if err != nil {
		log.Fatal(err)
	}
	data, err := sdr.Extract(csv, haders)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data {
		err := conn.Insert(d)
		if err != nil {
			log.Fatal(err)
		}
	}
}
