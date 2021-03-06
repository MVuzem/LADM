package surveying

import "github.com/MVuzem/LADM/shared"

// LAPoint Point
type LAPoint struct {
	shared.VersionedObject

	InterpolationRole LAInterpolationType
	Monumentation     *LAMonumentationType
	OriginalLocation  GMPoint
	PID               shared.Oid
	PointType         LAPointType
	ProductionMethod  *LILinage
	TransAndResult    *LATransformation
}

// GMPoint Point
type GMPoint string // TODO external package

// LILinage Linage
type LILinage string // TODO external package

// LAInterpolationType Interpolation role type
type LAInterpolationType int

const (
	// DefaultInterpolation Default Interpolation type
	DefaultInterpolation LAInterpolationType = 0
)

// LAMonumentationType Monumentation type
type LAMonumentationType int

const (
	// DefaultMonumentation Default Monumentation type
	DefaultMonumentation LAMonumentationType = 0
)

// LAPointType Point type
type LAPointType int

const (
	// DefaultPoint Default point type
	DefaultPoint LAPointType = 0
)
