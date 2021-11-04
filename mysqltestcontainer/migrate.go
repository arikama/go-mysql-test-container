package mysqltestcontainer

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/arikama/go-mysql-test-container/util"
)

func migrate(db *sql.DB, migrationDir string) error {
	fmt.Printf("🤖 Running migration...\n")
	migrationFiles, err := util.GetFiles(migrationDir)
	if err != nil {
		return err
	}
	sort.Sort(byFileVersion(migrationFiles))
	for i, migrationFile := range migrationFiles {
		fmt.Printf("🤖 #%v\t%v\n", i+1, migrationFile)
		content, _ := util.LoadFile(migrationFile)
		_, err = db.Exec(content)
		if err != nil {
			return err
		}
	}
	return nil
}
