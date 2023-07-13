package strategy

import (
	"fmt"
)

type StorageStrategy interface {
	Save(name string, data []byte) string
}

var strategies = map[string]StorageStrategy{
	"file": &fileStorageStrategy{},
	"db":   &dbStorageStrategy{},
}

func NewStorageStrategy(name string) (StorageStrategy, error) {
	if strategy, ok := strategies[name]; ok {
		return strategy, nil
	}
	return nil, fmt.Errorf("strategy %s not found", name)
}

type fileStorageStrategy struct{}

func (fs *fileStorageStrategy) Save(name string, data []byte) string {
	// save data to file
	return fmt.Sprintf("%s:%s saved data to file", name, data)
}

type dbStorageStrategy struct{}

func (ds *dbStorageStrategy) Save(name string, data []byte) string {
	// save data to db
	return fmt.Sprintf("%s:%s saved data to db", name, data)
}
