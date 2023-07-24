package dbmanager

import (
	"github.com/brownlow2/pdb/internal/db"
)

var (
	dbExistsError   = "database '%s' already exists"
	dbNotExistError = "database '%s' does not exist"
)

// DBManager is the interface for any DB manager instances
type DBManager interface {
	// Returns the map of DB instance names to their respective DB instance
	GetDBs() map[string]db.DB

	// Creates a DB instance and adds it to the DB map
	// Returns an error if the DB already exists
	// Returns an error if the keyHeader is empty
	// Returns an error if the list of headers is empty or the keyHeader is not in the list
	CreateDB(name string, headers []db.HeaderI, keyHeader string) error

	// Returns true if the DB exists in the DBManager's map of DBs
	DBExists(name string) bool
}

// The implementation for DBManager holding the following fields:
// DBs: the map containing the name of the DB mapped to the DB instance
type DBManagerImpl struct {
	DBs map[string]db.DB
}
