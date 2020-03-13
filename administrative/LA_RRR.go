package administrative

import "github.com/MVuzem/LADM/shared"

// LARRR right, restriction, responsibility
type LARRR struct {
	shared.VersionedObject

	RID         shared.Oid
	Description string
	Share       shared.Fraction
}
