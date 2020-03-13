package party

import "github.com/MVuzem/LADM/shared"

// LAParty LA Party
type LAParty struct {
	shared.VersionedObject

	// LA_Party
	ExtPid string
	Name   string
	Pid    string
	Role   string
	Type   string
}
