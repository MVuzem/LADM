package party

import "github.com/MVuzem/LADM/shared"

// LAPartyMember Party member
type LAPartyMember struct {
	shared.VersionedObject

	Share shared.Fraction
}
