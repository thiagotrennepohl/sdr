package sdr

//SdrRepository something
type SdrRepository interface {
	StoreData(interface{}) error
	StoreBatch([]interface{}) error
}
