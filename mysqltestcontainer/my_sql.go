package mysqltestcontainer

import "database/sql"

type MySql struct {
	Db       *sql.DB
	Username string
	Password string
	Ip       string
	Port     string
	Database string
}
