package surveying

import "github.com/MVuzem/LADM/shared"

// LABoundaryFaceString Boundary face string
type LABoundaryFaceString struct {
	shared.VersionedObject

	BfsID          shared.Oid
	Geometry       *GMMultiCurve
	LocationByText *string
}

// GMMultiCurve Multi curve
type GMMultiCurve string // TODO external package
