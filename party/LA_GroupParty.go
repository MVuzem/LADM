package party

import "github.com/MVuzem/LADM/shared"

// LAGroupParty Group party
type LAGroupParty struct {
	shared.VersionedObject

	// Versioned object
	ID                   string
	BeginLifespanVersion string
	EndLifespanVersion   string
}
