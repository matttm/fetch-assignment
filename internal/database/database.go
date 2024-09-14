package database

import (
	"fetch-assignment/internal/models"
	"fmt"
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
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
