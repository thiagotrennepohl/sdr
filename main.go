package main

import (
	sdrcli "github.com/thiagotrennepohl/sdr/sdr/cmd"
	repository "github.com/thiagotrennepohl/sdr/sdr/repository"
	usecase "github.com/thiagotrennepohl/sdr/sdr/usecase"
	mgo "gopkg.in/mgo.v2"
)

var mongoSession mgo.Session

func main() {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/somecollection")
	if err != nil {
		panic(err)
	}

	sdrRepository := repository.NewSdrRepository(mongoSession)
	sdrUseCase := usecase.NewSdrRepository(sdrRepository)

	sdrCli := sdrcli.NewSdrCli(sdrUseCase)

	csv, err := sdrCli.ReadCSV("dat.csv")
	if err != nil {
		panic(err)
	}
	headers, err := sdrCli.ParseHeaders(csv)
	if err != nil {
		panic(err)
	}

	err = sdrCli.StoreData(csv, headers)
}
