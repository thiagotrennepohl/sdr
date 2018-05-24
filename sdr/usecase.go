package sdr

type SdrUseCase interface {
	StoreData(map[string]interface{}) error
	StoreBatch([]interface{}) error
	ParseHeaders([]string) []string
}
