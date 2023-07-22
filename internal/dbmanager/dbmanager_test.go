package dbmanager

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brownlow2/pdb/internal/db"
)

func TestNew(t *testing.T) {
	dbm := New()
	assert.NotNil(t, dbm)
	assert.Equal(t, 0, len(dbm.DBs))
}

func testCreateDB(t *testing.T, name string, headers []string, errOccured bool) {
	dbm := &DBManagerImpl{
		DBs: map[string]db.DB{
			"existing db": &db.DBImpl{},
		},
	}

	err := dbm.CreateDB(name, headers)

	assert.Equal(t, errOccured, err != nil)
}

func TestCreateDB(t *testing.T) {
	testCreateDB(t, "new db", []string{}, false)
	testCreateDB(t, "existing db", []string{}, true)
}
