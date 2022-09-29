package mysqltestcontainer

import (
	"database/sql"
	"github.com/testcontainers/testcontainers-go"
)

type MySqlTestContainer struct {
	db        *sql.DB
	dbInfo    *DbInfo
	container testcontainers.Container
}

func (m *MySqlTestContainer) GetDb() *sql.DB {
	return m.db
}

func (m *MySqlTestContainer) GetDbInfo() *DbInfo {
	return m.dbInfo
}

func (m *MySqlTestContainer) GetContainer() testcontainers.Container {
	return m.container
}
