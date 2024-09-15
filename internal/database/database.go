package database

import (
	"fetch-assignment/internal/models"
	"sync"
)

type Database struct {
	TxTable map[string]*models.Transaction
}

var lock = &sync.Mutex{}

var instance *Database

func GetInstance() *Database {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Database{}
			instance.TxTable = make(map[string]*models.Transaction)
		}
	}

	return instance
}
