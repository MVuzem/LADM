package party

import "github.com/MVuzem/LADM/shared"

// LAGroupParty Group party
type LAGroupParty struct {
	LAParty

	GroupID shared.Oid
	Type    string
}
