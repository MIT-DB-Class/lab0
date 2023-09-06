package ridershipDB

type RidershipDB interface {
	Open(filePath string) error
	GetRidership(lineId string) ([]int64, error)
	Close() error
}
