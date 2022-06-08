package mysqltestcontainer

import "database/sql"

type MySql struct {
	db     *sql.DB
	dbInfo *DbInfo
}

func (m *MySql) GetDb() *sql.DB {
	return m.db
}

func (m *MySql) GetDbInfo() *DbInfo {
	return m.dbInfo
}
