package mysqltestcontainer

import (
	"database/sql"
	"fmt"
)

func open(ip, port, rootPassword, databaseName string) (*sql.DB, error) {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	dataSourceName := "root:" + rootPassword + "@tcp(" + ip + ":" + port + ")/" + databaseName + "?parseTime=true"

	fmt.Printf("🤖 DSN: %v\n", dataSourceName)
	return sql.Open("mysql", dataSourceName)
}
