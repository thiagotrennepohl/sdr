package usecase

import (
	"strings"

	"github.com/thiagotrennepohl/sdr/sdr"
)

type sdrUsecase struct {
	sdrRepository sdr.SdrRepository
}

func NewSdrRepository(sdrRepo sdr.SdrRepository) sdr.SdrUseCase {
	return &sdrUsecase{
		sdrRepository: sdrRepo,
	}
}

// func (sdr *sdrUsecase) StoreData(document interface{}) error {
// 	err := sdr.sdrRepository.StoreData(document)
// 	return err
// }

func (sdr *sdrUsecase) ParseHeaders(headers []string) []string {
	for index, value := range headers {
		value = strings.Trim(value, " ")
		value = strings.Replace(value, " ", "_", -1)
		headers[index] = strings.ToUpper(value)
	}
	return headers
}

func (sdr *sdrUsecase) StoreData(batch map[string]interface{}) error {

	return nil
}

func (sdr *sdrUsecase) StoreBatch(batch []map[string]interface{}) error {

}
