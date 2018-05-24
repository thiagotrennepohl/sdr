package repository

import (
	"github.com/thiagotrennepohl/sdr/sdr"
	"gopkg.in/mgo.v2"
)

type sdrRepository struct {
	session *mgo.Session
}

func NewSdrRepository(session *mgo.Session) sdr.SdrRepository {
	return &sdrRepository{
		session: session,
	}
}

func (sdr *sdrRepository) StoreData(document interface{}) error {
	conn := sdr.session.DB("").C("somecollection")
	err := conn.Insert(document)
	return err
}
