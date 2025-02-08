package system

import (
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"sync"
)

var lock = &sync.Mutex{}

type SystemParams struct {
	Dbc *sqldb.DBConnData
}

var sysparmInst *SystemParams

func GetSystemParams() *SystemParams {
	if sysparmInst == nil {
		lock.Lock()
		defer lock.Unlock()
		if sysparmInst == nil {
			sysparmInst = &SystemParams{}
		}
	}
	return sysparmInst
}

func (s *SystemParams) Close() {
	s.Dbc.Close()
}
