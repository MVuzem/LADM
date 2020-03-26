package spatialunit

import "github.com/MVuzem/LADM/shared"

// LASpatialUnitGroup Spatial unit group
type LASpatialUnitGroup struct {
	SugID          shared.Oid
	HierarchyLevel int
	Label          string
	Name           string
	ReferencePoint string
}
