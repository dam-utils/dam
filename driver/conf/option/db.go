package option

import "dam/config"

type DB struct {}

func (o *DB) GetType() string {
	return config.DB_TYPE
}
