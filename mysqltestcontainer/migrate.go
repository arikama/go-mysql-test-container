package mysqltestcontainer

import (
	"database/sql"
	"sort"
	"strings"

	"github.com/arikama/go-mysql-test-container/util"
	"github.com/hooligram/kifu"
)

func migrate(db *sql.DB, migrationDir string) error {
	kifu.Info("Running migration...")
	migrationFiles, err := util.GetFiles(migrationDir)
	if err != nil {
		return err
	}
	sort.Sort(byFileVersion(migrationFiles))
	for i, migrationFile := range migrationFiles {
		kifu.Info("%v\t%v", i+1, migrationFile)
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
