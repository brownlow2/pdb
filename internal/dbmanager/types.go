package dbmanager

type DBManager interface {
	CreateDB(name string, headers ...string) error
}
