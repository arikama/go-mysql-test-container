package mysqltestcontainer

import (
	"database/sql"
)

func open(ip, port, rootPassword, databaseName string) (*sql.DB, error) {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	dataSourceName := "root:" + rootPassword + "@tcp(" + ip + ":" + port + ")/" + databaseName + "?parseTime=true"

	return sql.Open("mysql", dataSourceName)
}
