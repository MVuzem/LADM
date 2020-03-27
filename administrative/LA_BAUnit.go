package administrative

import "github.com/MVuzem/LADM/shared"

// LABaunit Basic administrative unit
type LABaunit struct {
	shared.VersionedObject

	Name string
	Type LABAUnitType
	UID  shared.Oid
}

// LABAUnitType BA unit type
type LABAUnitType int

const (
	// DefaultBAUnit Default BA unit type
	DefaultBAUnit LABAUnitType = 0
)
