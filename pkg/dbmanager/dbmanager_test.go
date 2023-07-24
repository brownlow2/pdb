package dbmanager

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brownlow2/pdb/internal/db"
)

func TestNew(t *testing.T) {
	dbm := New()
	assert.NotNil(t, dbm)
	assert.Equal(t, 0, len(dbm.DBs))
}

func TestGetDBs(t *testing.T) {
	dbm := &DBManagerImpl{
		DBs: map[string]db.DB{
			"existing db": &db.DBImpl{},
		},
	}

	assert.True(t, reflect.DeepEqual(dbm.DBs, dbm.GetDBs()))
}

func testCreateDB(t *testing.T, name string, headers []db.HeaderI, keyHeader string, errOccured bool) {
	dbm := &DBManagerImpl{
		DBs: map[string]db.DB{
			"existing db": &db.DBImpl{},
		},
	}

	err := dbm.CreateDB(name, headers, keyHeader)

	assert.Equal(t, errOccured, err != nil)
}

func TestCreateDB(t *testing.T) {
	testCreateDB(t, "new db", []db.HeaderI{&db.Header{"Title", true, db.VALUE_STRING}}, "Title", false)
	testCreateDB(t, "existing db", []db.HeaderI{}, "", true)
	testCreateDB(t, "new db", []db.HeaderI{&db.Header{"Title", true, db.VALUE_STRING}}, "", true)
	testCreateDB(t, "new db", []db.HeaderI{}, "Title", true)
	testCreateDB(t, "new db", []db.HeaderI{&db.Header{"NotKey", true, db.VALUE_STRING}}, "Title", true)
}

func TestDBExists(t *testing.T) {
	dbi, err := db.New("test", []db.HeaderI{&db.Header{"Title", true, db.VALUE_STRING}}, "Title")
	assert.Nil(t, err)
	dbm := &DBManagerImpl{
		DBs: map[string]db.DB{
			"test": dbi,
		},
	}

	assert.True(t, dbm.DBExists("test"))
	assert.False(t, dbm.DBExists("Not exists"))
}
