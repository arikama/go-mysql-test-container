package mysqltestcontainer

import "github.com/arikama/go-mysql-test-container/util"

type byFileVersion []string

func (a byFileVersion) Len() int {
	return len(a)
}

func (a byFileVersion) Less(i, j int) bool {
	ai := util.GetVersion(a[i])
	aj := util.GetVersion(a[j])
	return ai < aj
}

func (a byFileVersion) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
