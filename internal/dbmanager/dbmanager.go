package dbmanager

import (
	"errors"
	"fmt"

	"github.com/brownlow2/pdb/internal/db"
)

type DBManagerImpl struct {
	DBs map[string]db.DB
}

func New() *DBManagerImpl {
	return &DBManagerImpl{
		DBs: map[string]db.DB{},
	}
}

func (dbm *DBManagerImpl) CreateDB(name string, headers []string, keyHeader string) error {
	if _, exists := dbm.DBs[name]; exists {
		dbExists := fmt.Sprintf("database '%s' already exists", name)
		return errors.New(dbExists)
	}

	db, err := db.New(name, headers, keyHeader)
	if err != nil {
		return err
	}

	dbm.DBs[name] = db

	return nil
}
