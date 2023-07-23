package dbmanager

import (
	"errors"
	"fmt"

	"github.com/brownlow2/pdb/internal/db"
)

func New() *DBManagerImpl {
	return &DBManagerImpl{
		DBs: map[string]db.DB{},
	}
}

func (dbm *DBManagerImpl) GetDBs() map[string]db.DB {
	return dbm.DBs
}

func (dbm *DBManagerImpl) CreateDB(name string, headers []db.HeaderI, keyHeader string) error {
	if dbm.DBExists(name) {
		return errors.New(fmt.Sprintf(dbExistsError, name))
	}

	db, err := db.New(name, headers, keyHeader)
	if err != nil {
		return err
	}

	dbm.DBs[name] = db

	return nil
}

func (dbm *DBManagerImpl) DBExists(name string) bool {
	_, exists := dbm.DBs[name]
	return exists
}
