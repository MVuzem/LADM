package party

import "github.com/MVuzem/LADM/shared"

// LAParty LA Party
type LAParty struct {
	shared.VersionedObject

	// LA_Party
	ExtPid *shared.Oid
	Name   string
	Pid    shared.Oid
	Role   string
	Type   string
}
