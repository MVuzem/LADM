package administrative

import "github.com/MVuzem/LADM/shared"

// LABaunit Basic administrative unit
type LABaunit struct {
	shared.VersionedObject

	Name string
	Type string
	UID  shared.Oid
}
