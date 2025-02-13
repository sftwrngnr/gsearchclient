package system

import (
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"sync"
)

var lock = &sync.Mutex{}

type SystemParams struct {
	Dbc   *sqldb.DBConnData
	GHost string
	GQKey string
	pLock *sync.Mutex
}

var sysparmInst *SystemParams

func GetSystemParams() *SystemParams {
	if sysparmInst == nil {
		lock.Lock()
		defer lock.Unlock()
		if sysparmInst == nil {
			sysparmInst = &SystemParams{pLock: &sync.Mutex{}}

		}
	}
	return sysparmInst
}

func (s *SystemParams) Close() {
	s.Dbc.Close()
}

func (SP *SystemParams) ParmLock() {
	SP.pLock.Lock()
}

func (SP *SystemParams) ParmUnlock() {
	SP.pLock.Unlock()
}
