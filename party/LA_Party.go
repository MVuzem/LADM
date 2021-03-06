package party

import "github.com/MVuzem/LADM/shared"

// LAParty LA Party
type LAParty struct {
	shared.VersionedObject

	// LA_Party
	ExtPid  *shared.Oid
	Name    *string
	Pid     shared.Oid
	Role    *LAPartyRoleType
	Type    LAPartyType
	Members *LAPartyMember
}

// LAPartyType Party type
type LAPartyType int

const (
	// DefaultParty Default party type
	DefaultParty LAPartyType = 0
)

// LAPartyRoleType Party role type
type LAPartyRoleType int

const (
	// DefaultPartyRole Default Party role type
	DefaultPartyRole LAPartyRoleType = 0
)
