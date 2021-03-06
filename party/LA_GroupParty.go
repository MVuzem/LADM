package party

import "github.com/MVuzem/LADM/shared"

// LAGroupParty Group party
type LAGroupParty struct {
	LAParty

	GroupID shared.Oid
	Type    LAGroupPartyType
	Members *LAPartyMember
}

// LAGroupPartyType Group party type
type LAGroupPartyType int

const (
	// DefaultGroupParty Default Group party type
	DefaultGroupParty LAGroupPartyType = 0
)
