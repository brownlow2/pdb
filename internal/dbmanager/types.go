package dbmanager

import (
	"github.com/brownlow2/pdb/internal/db"
)

var (
	dbExistsError = "database '%s' already exists"
)

type DBManager interface {
	GetDBs() map[string]db.DB
	CreateDB(name string, headers []db.HeaderI, keyHeader string) error
	DBExists(name string) bool
}

type DBManagerImpl struct {
	DBs map[string]db.DB
}
