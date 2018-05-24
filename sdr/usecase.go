package sdr

type SdrUseCase interface {
	StoreData(map[string]interface{}) error
	StoreBatch([]map[string]interface{}) error
	ParseHeaders([]string) []string
}
