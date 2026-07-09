package repo

type Database interface{}

type DatabaseImpl struct {
}

func NewDatabase() Database {
	return DatabaseImpl{}
}
