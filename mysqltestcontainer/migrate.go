package mysqltestcontainer

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

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
		statements := strings.Split(content, ";")
		for _, statement := range statements {
			if statement == "" {
				continue
			}
			_, err = db.Exec(statement + ";")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
